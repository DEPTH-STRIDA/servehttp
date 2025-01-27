package service

import (
	"pet1/internal/model"
	"pet1/internal/repo"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(user model.User) (model.User, error) {
	return s.repo.CreateUser(user)
}

// GetAllUsers возвращает всех пользователей
func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

// UpdateUserByID обновляет пользователя по ID
func (s *UserService) UpdateUserByID(id uint, user model.User) (model.User, error) {
	return s.repo.UpdateUserByID(id, user)
}

// DeleteUserByID удаляет пользователя по ID
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

// GetTasksForUser получает все задачи пользователя
func (s *UserService) GetTasksForUser(userID uint) ([]model.Task, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user.Tasks, nil
}
