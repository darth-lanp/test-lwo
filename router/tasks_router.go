package router

import (
	"net/http"
	"rest/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(taskController *controller.TaskController) {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/tag")
	tagRouter.GET("", taskController.FindAll)
	tagRouter.GET("/:tagId", taskController.FindById)
	tagRouter.POST("", taskController.Create)
	tagRouter.PATCH("/:tagId", taskController.Update)
	tagRouter.DELETE("/:tagId", taskController.Delete)

	return service
}
