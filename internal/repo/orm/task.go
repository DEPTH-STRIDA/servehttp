package orm

import (
	"errors"
	"pet1/internal/model"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task model.Task) (model.Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return model.Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// UpdateTaskByID обновляет задачу по ее ID
func (r *taskRepository) UpdateTaskByID(id uint, task model.Task) (model.Task, error) {
	// Ищем задачу в базе данных по ID
	var existingTask model.Task
	result := r.db.First(&existingTask, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Если задача не найдена, возвращаем ошибку
			return model.Task{}, errors.New("task not found")
		}
		// Если возникла другая ошибка при поиске, возвращаем её
		return model.Task{}, result.Error
	}

	// Обновляем поля задачи, если они предоставлены
	if task.Task != "" {
		existingTask.Task = task.Task
	}
	existingTask.IsDone = task.IsDone

	// Сохраняем обновленную задачу в базе данных
	saveResult := r.db.Save(&existingTask)
	if saveResult.Error != nil {
		return model.Task{}, saveResult.Error
	}

	// Возвращаем обновленную задачу и nil (отсутствие ошибки)
	return existingTask, nil
}

// DeleteTaskByID удаляет задачу по ее ID
func (r *taskRepository) DeleteTaskByID(id uint) error {
	// Ищем задачу в базе данных по ID
	var task model.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Если задача не найдена, возвращаем ошибку
			return errors.New("task not found")
		}
		// Если возникла другая ошибка при поиске, возвращаем её
		return result.Error
	}

	// Удаляем задачу из базы данных
	deleteResult := r.db.Delete(&task)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	// Проверяем, была ли удалена хотя бы одна запись
	if deleteResult.RowsAffected == 0 {
		return errors.New("no task was deleted")
	}

	// Возвращаем nil, указывая на отсутствие ошибки
	return nil
}

// GetTasksByUserID получает все задачи пользователя по его ID
func (r *taskRepository) GetTasksByUserID(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
