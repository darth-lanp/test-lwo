package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"rest/config"
	"rest/controller"
	"rest/repository"
	"rest/router"
	"rest/service"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db := config.DatabaseConnection()
	defer db.Close()

	validate := validator.New()

	tagRepository := repository.NewTagsRepositoryImpl(db)

	tagService := service.NewTagsServiceImpl(tagRepository, validate)

	tagController := controller.NewTagController(tagService)

	router := router.NewRouter(tagController)

	signalChan := make(chan os.Signal, 1)
	stopChan := make(chan struct{}, 1)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case <-signalChan:
				slog.Info("Stop fon task")
				stopChan <- struct{}{}
				return
			case <-ticker.C:
				tasks, _ := tagRepository.FindAll()
				for _, task := range tasks {
					if task.DueDate.Before(time.Now()) && !task.Overdue {
						tagRepository.OverdueTask(task.Id)
						slog.Info("Overdue task with", slog.Int("id", task.Id))
					}
				}
			}
		}
	}()

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(router),
	}
	go func() {
		slog.Info("Run server on port", slog.Int("port", 8080))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server error: ", slog.String("err", err.Error()))
		}
	}()

	<-stopChan
	err := server.Shutdown(context.Background())
	if err != nil {
		slog.Error("Error shutdown server: ", slog.String("err", err.Error()))
	} else {
		slog.Info("Server stop")
	}
}
