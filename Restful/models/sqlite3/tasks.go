package sqlite3

import (
	"database/sql"
	"os"
	"time"
)

type Task struct {
	Id          int
	Name        string
	Data        string
	CreatedDate time.Time
}

type TasksModel struct {
	Db *sql.DB
}

func (model *TasksModel) AddTask(task *Task) (*Task, error) {
	stmt := "INSERT INTO tasks (name, data, createdDate) VALUES (?, ?, datetime('now'))"
	if task.Name == "" || task.Data == "" {
		return nil, os.ErrInvalid
	}
	data, err := model.Db.Exec(stmt, task.Name, task.Data)
	if err != nil {
		return nil, err
	}
	insertId, err := data.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.Id = int(insertId)
	return task, nil
}

func (model *TasksModel) GetById(id uint) (*Task, error) {
	stmt := "SELECT id, name, data, createdDate FROM tasks WHERE id=(?)"
	row := model.Db.QueryRow(stmt, id)
	task := Task{}
	err := row.Scan(&task.Id, &task.Name, &task.Data, &task.CreatedDate)
	if err != nil {
		return nil, err
	}
	return &task, nil

}

func (model *TasksModel) DeleteById(id uint) error {
	stmt := "DELETE FROM tasks WHERE id=(?)"
	_, err := model.Db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (model *TasksModel) GetAll() ([]Task, error) {
	stmt := "SELECT id, name, data, createdDate FROM tasks"
	rows, err := model.Db.Query(stmt)
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0)
	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.Id, &task.Name, &task.Data, &task.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (model *TasksModel) DeleteAll() (int, error) {
	stmt := "DELETE FROM tasks"
	ret, err := model.Db.Exec(stmt)
	if err != nil {
		return 0, err
	}
	count, err := ret.RowsAffected()
	if err != nil {
		return int(count), err
	}
	return int(count), nil

}
