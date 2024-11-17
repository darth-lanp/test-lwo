package service

import (
	"fmt"
	"rest/data/request"
	"rest/data/response"
	"rest/model"
	"rest/repository"
	"time"

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

func (t TaskServiceImpl) Create(task request.CreateTaskRequest) (response.TaskResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service create: %w", err)
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, task.DueDate)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service create: %w", err)
	}

	taskModel := model.Task{
		Title:       task.Title,
		Description: task.Description,
		DueDate:     date,
		Overdue:     false,
		Completed:   false,
	}

	newTask, err := t.TagRepository.Save(taskModel)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service create: %w", err)
	}
	return response.TaskResponse(newTask), nil
}

func (t TaskServiceImpl) Update(task request.UpdateTaskRequest) (response.TaskResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service update: %w", err)
	}

	taskData, err := t.TagRepository.FindById(task.Id)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service update: %w", err)
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, task.DueDate)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service update: %w", err)
	}

	taskData.Title = task.Title
	taskData.Description = task.Description
	taskData.DueDate = date

	err = t.TagRepository.Update(taskData)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service update: %w", err)
	}
	return response.TaskResponse(taskData), nil
}

func (t TaskServiceImpl) CompletedTask(taskId int) (response.TaskResponse, error) {
	taskData, err := t.TagRepository.FindById(taskId)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service update: %w", err)
	}

	err = t.TagRepository.CompletedTask(taskId)
	if err != nil {
		return response.TaskResponse{}, fmt.Errorf("service completed: %w", err)
	}
	taskData.Completed = true
	return response.TaskResponse(taskData), nil
}

func (t TaskServiceImpl) Delete(taskId int) error {
	_, err := t.TagRepository.FindById(taskId)
	if err != nil {
		return fmt.Errorf("service update: %w", err)
	}

	err = t.TagRepository.Delete(taskId)
	if err != nil {
		return fmt.Errorf("service delete: %w", err)
	}
	return nil
}

func (t TaskServiceImpl) FindById(taskId int) (response.TaskResponse, error) {
	taskData, err := t.TagRepository.FindById(taskId)
	if err != nil {
		return response.TaskResponse{}, err
	}

	taskResponse := response.TaskResponse{
		Id:          taskData.Id,
		Title:       taskData.Title,
		Description: taskData.Description,
		DueDate:     taskData.DueDate,
	}

	return taskResponse, nil
}

func (t TaskServiceImpl) FindAll() ([]response.TaskResponse, error) {
	result, err := t.TagRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("service findall: %w", err)
	}
	var tags []response.TaskResponse
	for _, value := range result {
		tags = append(tags, response.TaskResponse{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			DueDate:     value.DueDate,
			Overdue:     value.Overdue,
			Completed:   value.Completed,
		})
	}
	return tags, err
}
