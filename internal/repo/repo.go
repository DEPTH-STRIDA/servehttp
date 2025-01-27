package repo

import "pet1/internal/model"

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task model.Task) (model.Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]model.Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task model.Task) (model.Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userID uint) ([]model.Task, error)
}

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
	UpdateUserByID(id uint, user model.User) (model.User, error)
	DeleteUserByID(id uint) error
}
