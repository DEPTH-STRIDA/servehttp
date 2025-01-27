package orm

import (
	"errors"
	"pet1/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (model.User, error) {
	var user model.User
	result := r.db.Preload("Tasks").First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUserByID(id uint, user model.User) (model.User, error) {
	var existingUser model.User
	result := r.db.First(&existingUser, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, result.Error
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}

	saveResult := r.db.Save(&existingUser)
	if saveResult.Error != nil {
		return model.User{}, saveResult.Error
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}

	deleteResult := r.db.Delete(&user)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	if deleteResult.RowsAffected == 0 {
		return errors.New("no user was deleted")
	}

	return nil
}
