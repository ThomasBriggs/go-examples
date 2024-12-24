package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		panic("Could not open database!")
	}
	defer db.Close()

	r := gin.Default()
	r.StaticFile("/", "./static/index.html")

	SetupTaskRoutes(r, db)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
