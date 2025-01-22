package handlers

import (
	"context"
	"fmt"
	"pet1/internal/taskService"
	"pet1/internal/web/tasks"
)

// Handler нужна для создания структуры  на этапе инициализации приложения
type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// DeleteTasksId реализует удаление задачи по ID
func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Извлекаем ID задачи из запроса
	id := request.Id

	// Вызываем сервис для удаления задачи
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		if err.Error() == "task not found" || err.Error() == "no task was deleted" {
			// Возвращаем 404 Not Found, если задача не найдена
			return tasks.DeleteTasksId404Response{}, nil
		}
		// Возвращаем 500 Internal Server Error для других ошибок
		return nil, fmt.Errorf("failed to delete task: %w", err)
	}

	// Возвращаем 204 No Content при успешном удалении
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Извлекаем ID задачи из запроса
	id := request.Id

	// Извлекаем тело запроса, содержащее поля для обновления
	body := request.Body

	// Создаём объект задачи с обновлёнными полями
	taskToUpdate := taskService.Task{}

	if body.Task != nil {
		taskToUpdate.Task = *body.Task
	}

	if body.IsDone != nil {
		taskToUpdate.IsDone = *body.IsDone
	}

	// Вызываем сервис для обновления задачи
	updatedTask, err := h.Service.UpdateTaskByID(id, taskToUpdate)
	if err != nil {
		if err.Error() == "task not found" {
			// Возвращаем 404 Not Found, если задача не найдена
			return tasks.PatchTasksId404Response{}, nil
		}
		// Возвращаем 500 Internal Server Error для других ошибок
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	// Создаём объект ответа с обновлённой задачей
	responseTask := tasks.Task{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем 200 OK с обновлённой задачей
	return tasks.PatchTasksId200JSONResponse(responseTask), nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}
