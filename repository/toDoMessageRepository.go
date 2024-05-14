package repository

import (
	"sync"

	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoMessageArray struct {
	toDoMessages []models.ToDoMessage
	mutex        sync.Mutex
}

type ToDoMessageRepository interface {
	CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error)
	DeleteToDoMessageByMessageID(messageID, userID uint) error
	UpdateToDoMessageByMessageID(messageID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error)
}
