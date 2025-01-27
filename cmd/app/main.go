package main

import (
	"log"
	"pet1/internal/api/handlers"
	"pet1/internal/api/web/tasks"
	"pet1/internal/api/web/users"
	"pet1/internal/db"
	"pet1/internal/repo/orm"
	"pet1/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация БД
	db.InitDB()

	// Инициализация сервисов задач
	tasksRepo := orm.NewTaskRepository(db.DB)
	tasksService := service.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	// Инициализация сервисов пользователей
	usersRepo := orm.NewUserRepository(db.DB)
	usersService := service.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация обработчиков задач
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Регистрация обработчиков пользователей
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
