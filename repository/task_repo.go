package repository

import "rest/model"

type TaskRepository interface {
	Save(task model.Task) model.Task
	Update(task model.Task) model.Task
	Delete(taskId int) error
	FindById(taskId int) (task model.Task, err error)
	FindAll() []model.Task
}
