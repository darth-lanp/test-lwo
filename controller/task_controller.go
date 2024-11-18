package controller

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"regexp"
	"rest/data/request"
	"rest/repository"
	"rest/service"
	"strconv"
)

type TaskController struct {
	tagService service.TaskService
}

func NewTaskController(tagService service.TaskService) *TaskController {
	return &TaskController{
		tagService: tagService,
	}
}

func (controller *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	createTagRequest := request.CreateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&createTagRequest)
	if err != nil {
		slog.Error("Error decode body", slog.String("err", err.Error()))
		http.Error(w, "Error create task", http.StatusInternalServerError)
		return
	}

	newTask, err := controller.tagService.Create(createTagRequest)
	if err != nil {
		slog.Error("Error create task", slog.String("err", err.Error()), slog.Any("body", createTagRequest))
		http.Error(w, "Error create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		panic(err)
	}
	slog.Info("Succsesfully create task", slog.Any("body", newTask))
}

func (controller *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	updateTagRequest := request.UpdateTaskRequest{}

	taskId := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(taskId)
	if err != nil {
		slog.Error("Invalid task ID", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updateTagRequest)
	if err != nil {
		http.Error(w, "Error decode body", http.StatusBadRequest)
		return
	}

	updateTagRequest.Id = id
	updateTask, err := controller.tagService.Update(updateTagRequest)
	if errors.Is(err, repository.NotFoundError) {
		slog.Error("Not found task", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Not found task", http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("Error update task", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Error update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updateTask)
	if err != nil {
		panic(err)
	}
	slog.Info("Succsesfully update task", slog.Any("body", updateTask))
}

func (controller *TaskController) CompletedTask(w http.ResponseWriter, r *http.Request) {
	re1 := regexp.MustCompile(`^/tasks/(\d+)/complete$`)
	mathes := re1.FindStringSubmatch(r.URL.Path)

	taskId := mathes[1]
	id, err := strconv.Atoi(taskId)
	if err != nil {
		slog.Error("Invalid task ID", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	completedTask, err := controller.tagService.CompletedTask(id)
	if errors.Is(err, repository.NotFoundError) {
		slog.Error("Not found task", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Not found task", http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("Error completed task", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Error completed task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(completedTask)
	if err != nil {
		panic(err)
	}
	slog.Info("Succsesfully completed task")
}

func (controller *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(taskId)
	if err != nil {
		slog.Error("Invalid task ID", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = controller.tagService.Delete(id)
	if errors.Is(err, repository.NotFoundError) {
		slog.Error("Not found task", slog.Int("id", id), slog.String("err", err.Error()))
		http.Error(w, "Not found task", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	slog.Info("Succsesfully deleted task", slog.Int("id", id))
}

func (controller *TaskController) FindAll(w http.ResponseWriter, r *http.Request) {
	tasksResponse, err := controller.tagService.FindAll()
	if err != nil {
		slog.Error("Error get all tasks", slog.String("err", err.Error()))
		http.Error(w, "Error get all tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasksResponse)
	if err != nil {
		panic(err)
	}
	slog.Info("Succsesfully get tasks")
}
