package service

import (
	"pet1/internal/model"
	"pet1/internal/repo"
)

type TaskService struct {
	repo repo.TaskRepository
}

func NewTaskService(repo repo.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task model.Task) (model.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task model.Task) (model.Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]model.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
