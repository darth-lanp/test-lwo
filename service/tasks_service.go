package service

import (
	"rest/data/request"
	"rest/data/response"
)

type TaskService interface {
	Create(task request.CreateTaskRequest) (response.TaskResponse, error)
	Update(tags request.UpdateTaskRequest) (response.TaskResponse, error)
	CompletedTask(tagsId int) (response.TaskResponse, error)
	Delete(tagsId int) error
	FindById(tagsId int) (response.TaskResponse, error)
	FindAll() ([]response.TaskResponse, error)
}
