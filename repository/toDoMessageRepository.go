package repository

import (
	"github.com/nazzarr03/To-Do-Rest-Api/config"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoMessageMockRepository struct{}

type ToDoMessageRepository interface {
	GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error)
}

func (toDoMessageMockRepo *ToDoMessageMockRepository) GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error) {
	var toDoMessages []models.ToDoMessage

	for _, toDoMessage := range config.ToDoMessages {
		if toDoMessage.DeletedAt.IsZero() && toDoMessage.ListID == listID {
			toDoMessages = append(toDoMessages, toDoMessage)
		}
	}

	return toDoMessages, nil
}
