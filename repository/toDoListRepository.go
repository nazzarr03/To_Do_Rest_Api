package repository

import (
	"github.com/nazzarr03/To-Do-Rest-Api/config"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoListMockRepository struct{}

type ToDoListRepository interface {
	GetTodoListsByUserID(userID uint) ([]models.ToDoList, error)
}

func (toDoListMockRepo *ToDoListMockRepository) GetTodoListsByUserID(userID uint) ([]models.ToDoList, error) {
	var todoLists []models.ToDoList

	for _, item := range config.ToDoLists {
		if item.UserID == userID && item.DeletedAt.IsZero() {
			todoLists = append(todoLists, item)
		}
	}

	return todoLists, nil
}
