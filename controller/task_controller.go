package controller

import (
	"net/http"
	"rest/data/request"
	"rest/data/response"
	"rest/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	tagService service.TaskService
}

func NewTagController(tagService service.TaskService) *TaskController {
	return &TaskController{
		tagService: tagService,
	}
}

func (controller *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	createTagRequest := request.CreateTaskRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	if err != nil {
		panic(err)
	}
	newTask := controller.tagService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   newTask,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TaskController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTaskRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	if err != nil {
		panic(err)
	}

	tagId := ctx.Param("taskId")
	id, err := strconv.Atoi(tagId)

	if err != nil {
		panic(err)
	}

	updateTagRequest.Id = id
	updateTask := controller.tagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   updateTask,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TaskController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("taskId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		panic(err)
	}

	var webResponse response.Response
	err = controller.tagService.Delete(id)
	if err != nil {
		webResponse = response.Response{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
		}
		ctx.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse = response.Response{
			Code:   200,
			Status: "Ok",
			Data:   nil,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *TaskController) FindById(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	id, err := strconv.Atoi(taskId)

	if err != nil {
		panic(err)
	}

	tasksResponse := controller.tagService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tasksResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TaskController) FindAll(ctx *gin.Context) {
	tasksResponse := controller.tagService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tasksResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
