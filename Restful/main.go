package main

import (
	"Restful/models/sqlite3"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		panic("Could not open database!")
	}
	defer db.Close()

	tasksModel := sqlite3.TasksModel{Db: db}

	r := gin.Default()

	// taskRouts := r.Group("/tasks")
	// {
	// 	taskRouts.GET("", )
	// }

	r.POST("/tasks", func(c *gin.Context) {
		var task sqlite3.Task
		c.BindJSON(&task)
		insertedTask, err := tasksModel.AddTask(&task)
		if err != nil {
			if errors.Is(err, os.ErrInvalid) {
				slog.Warn("Invalid arguments given!")
				c.Status(http.StatusBadRequest)
				return
			}
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, insertedTask)
	})

	r.GET("/tasks", func(c *gin.Context) {
		tasks, _ := tasksModel.GetAll()
		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:taskId", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("taskId"), 10, 0)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		task, err := tasksModel.GetById(uint(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.Status(http.StatusNotFound)
				return
			}
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.DELETE("/tasks", func(ctx *gin.Context) {
		count, err := tasksModel.DeleteAll()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, struct{ Count int }{count})
	})

	r.DELETE("/tasks/:taskId", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("taskId"), 10, 0)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = tasksModel.DeleteById(uint(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.Status(http.StatusNotFound)
				return
			}
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK)
	})

	r.StaticFile("/", "./static/index.html")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}