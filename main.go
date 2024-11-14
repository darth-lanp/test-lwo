package main

import (
	"net/http"
	"rest/config"
	"rest/controller"
	"rest/repository"
	"rest/router"
	"rest/service"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.DatabaseConnection()
	defer db.Close()

	validate := validator.New()

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagsServiceImpl(tagRepository, validate)

	//Init controller
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
