package service

import (
	"rest/data/request"
	"rest/data/response"
)

type TaskService interface {
	Create(task request.CreateTaskRequest) response.TaskResponse
	Update(tags request.UpdateTaskRequest) response.TaskResponse
	Delete(tagsId int) error
	FindById(tagsId int) response.TaskResponse
	FindAll() []response.TaskResponse
}
