package service

import (
	"rest/data/request"
	"rest/data/response"
	"rest/model"
	"rest/repository"

	"github.com/go-playground/validator/v10"
)

type TaskServiceImpl struct {
	TagRepository repository.TaskRepository
	Validate      *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TaskRepository, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}

func (t TaskServiceImpl) Create(task request.CreateTaskRequest) response.TaskResponse {
	err := t.Validate.Struct(task)
	if err != nil {
		panic(err)
	}

	taskModel := model.Task{
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
	}

	newTask := t.TagRepository.Save(taskModel)
	return response.TaskResponse(newTask)
}

func (t TaskServiceImpl) Update(task request.UpdateTaskRequest) response.TaskResponse {
	taskData, err := t.TagRepository.FindById(task.Id)
	if err != nil {
		panic(err)
	}

	taskData.Title = task.Title
	taskData.Description = task.Description
	taskData.DueDate = task.DueDate

	updateTask := t.TagRepository.Update(taskData)

	return response.TaskResponse(updateTask)
}

func (t TaskServiceImpl) Delete(taskId int) error {
	err := t.TagRepository.Delete(taskId)
	return err
}

func (t TaskServiceImpl) FindById(taskId int) response.TaskResponse {
	taskData, err := t.TagRepository.FindById(taskId)
	if err != nil {
		panic(err)
	}

	taskResponse := response.TaskResponse{
		Id:          taskData.Id,
		Title:       taskData.Title,
		Description: taskData.Description,
		DueDate:     taskData.DueDate,
	}

	return taskResponse
}

func (t TaskServiceImpl) FindAll() []response.TaskResponse {
	result := t.TagRepository.FindAll()
	var tags []response.TaskResponse
	for _, value := range result {
		tags = append(tags, response.TaskResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			DueDate:     value.DueDate,
		})
	}
	return tags
}
