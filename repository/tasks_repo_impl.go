package repository

import (
	"database/sql"
	"errors"
	"rest/model"
	"time"
)

type TaskRepositoryImpl struct {
	Db *sql.DB
}

func NewTagsRepositoryImpl(Db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: Db}
}

func (t TaskRepositoryImpl) Save(task model.Task) model.Task {
	result, err := t.Db.Exec(
		"insert into Tasks (title, description, duedate) values ($1, $2, $3)",
		task.Title,
		task.Description,
		task.DueDate,
	)
	if err != nil {
		panic(err)
	}

	newTaskId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return model.Task{
		Id:          int(newTaskId),
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
	}
}

func (t TaskRepositoryImpl) Delete(taskId int) error {
	_, err := t.Db.Exec("delete from Tasks wehre if = $1", taskId)
	return err
}

func (t TaskRepositoryImpl) Update(task model.Task) model.Task {
	// var updateTask = request.UpdateTaskRequest{
	// 	Id:          task.Id,
	// 	Title:       task.Title,
	// 	Description: task.Description,
	// 	DueDate:     task.DueDate,
	// }
	_, err := t.Db.Exec("update Tasks set title = $1, description = $2, due_date = $3 where id = $4",
		task.Title,
		task.Description,
		task.DueDate,
	)
	if err != nil {
		panic(err)
	}
	return task
}

func (t TaskRepositoryImpl) FindById(taskId int) (model.Task, error) {
	row := t.Db.QueryRow("select * from Tasks where id = $1", taskId)
	var c1 int
	var c2, c3 string
	var c4 time.Time
	err := row.Scan(&c1, &c2, &c3, &c4)
	if err != nil {
		return model.Task{}, errors.New("task is not fount")
	}
	task := model.Task{c1, c2, c3, c4}
	return task, nil
}

func (t TaskRepositoryImpl) FindAll() []model.Task {
	rows, err := t.Db.Query("select * from Tasks")
	if err != nil {
		panic(err)
	}

	var tasks []model.Task
	var c1 int
	var c2, c3 string
	var c4 time.Time
	for rows.Next() {
		_ = rows.Scan(&c1, &c2, &c3, &c4)
		task := model.Task{c1, c2, c3, c4}
		tasks = append(tasks, task)
	}
	return tasks
}
