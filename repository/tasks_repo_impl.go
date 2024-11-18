package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"rest/model"
	"time"
)

var NotFoundError = errors.New("not found object")

type TaskRepositoryImpl struct {
	Db *sql.DB
}

func NewTaskRepositoryImpl(Db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: Db}
}

func (t TaskRepositoryImpl) Save(task model.Task) (model.Task, error) {
	result, err := t.Db.Exec(
		"insert into Tasks (title, description, duedate, overdue, completed) values ($1, $2, $3, $4, $5)",
		task.Title,
		task.Description,
		task.DueDate,
		task.Overdue,
		task.Completed,
	)
	if err != nil {
		return model.Task{}, fmt.Errorf("repository save: %w", err)
	}

	newTaskId, err := result.LastInsertId()
	if err != nil {
		return model.Task{}, fmt.Errorf("repository save: %w", err)
	}

	return model.Task{
		Id:          int(newTaskId),
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Overdue:     task.Overdue,
		Completed:   task.Completed,
	}, nil
}

func (t TaskRepositoryImpl) Delete(taskId int) error {
	_, err := t.Db.Exec("delete from Tasks where id = $1", taskId)
	if err != nil {
		return fmt.Errorf("repository delete: %w", err)
	}
	return nil
}

func (t TaskRepositoryImpl) CompletedTask(taskId int) error {
	_, err := t.Db.Exec("update Tasks set completed = TRUE where id = $1", taskId)
	if err != nil {
		return fmt.Errorf("repository cpmpleted: %w", err)
	}
	return nil
}

func (t TaskRepositoryImpl) OverdueTask(taskId int) error {
	_, err := t.Db.Exec("update Tasks set overdue = TRUE where id = $1", taskId)
	if err != nil {
		return fmt.Errorf("overdue cpmpleted: %w", err)
	}
	return nil
}

func (t TaskRepositoryImpl) Update(task model.Task) error {
	_, err := t.Db.Exec("update Tasks set title = $1, description = $2, duedate = $3 where id = $4",
		task.Title,
		task.Description,
		task.DueDate,
		task.Id,
	)
	if err != nil {
		return fmt.Errorf("repository update: %w", err)
	}
	return err
}

func (t TaskRepositoryImpl) FindById(taskId int) (model.Task, error) {
	fmt.Println(taskId)
	row := t.Db.QueryRow("select * from Tasks where id = $1", taskId)
	var c1 int
	var c2, c3 string
	var c4 time.Time
	var c5, c6 bool
	err := row.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
	if err != nil {
		return model.Task{}, NotFoundError
	}
	task := model.Task{
		Id:          c1,
		Title:       c2,
		Description: c3,
		DueDate:     c4,
		Overdue:     c5,
		Completed:   c6,
	}
	return task, nil
}

func (t TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	rows, err := t.Db.Query("select * from Tasks")
	if err != nil {
		return nil, fmt.Errorf("repository findAll: %w", err)
	}
	var tasks []model.Task
	var c1 int
	var c2, c3 string
	var c4 time.Time
	var c5, c6 bool
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			return nil, fmt.Errorf("repository findAll: %w", err)
		}
		task := model.Task{
			Id:          c1,
			Title:       c2,
			Description: c3,
			DueDate:     c4,
			Overdue:     c5,
			Completed:   c6,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
