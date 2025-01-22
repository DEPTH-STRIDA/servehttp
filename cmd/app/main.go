package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet1/internal/db"
	"pet1/internal/handlers"
	"pet1/internal/taskService"
	"pet1/internal/web/tasks"
)

func main() {
	db.InitDB()
	if err := db.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatal(err)
	}

	repo := taskService.NewTaskRepository(db.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
