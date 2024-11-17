package repository

import "rest/model"

type TaskRepository interface {
	Save(task model.Task) (model.Task, error)
	Update(task model.Task) error
	CompletedTask(taskId int) error
	OverdueTask(taskId int) error
	Delete(taskId int) error
	FindById(taskId int) (task model.Task, err error)
	FindAll() ([]model.Task, error)
}
