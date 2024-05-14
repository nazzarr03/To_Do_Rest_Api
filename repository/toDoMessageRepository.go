package repository

import (
	"errors"
	"sync"
	"time"

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

func (todoMessageArray *ToDoMessageArray) CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error) {
	todoMessageArray.mutex.Lock()
	defer todoMessageArray.mutex.Unlock()

	toDoListArray := ToDoListArray{}
	userArray := UserArray{}

	for _, toDoList := range toDoListArray.toDoLists {
		if toDoList.ListID == listID {
			toDoMessage.ListID = listID
		} else {
			return models.ToDoMessage{}, errors.New("to-do list not found")
		}
	}

	for _, user := range userArray.users {
		if user.UserID == userID {
			toDoMessage.UserID = userID
		} else {
			return models.ToDoMessage{}, errors.New("user not found")
		}
	}

	toDoMessage.MessageID = uint(todoMessageArray.toDoMessages[len(todoMessageArray.toDoMessages)-1].MessageID) + 1
	toDoMessage.ListID = listID
	toDoMessage.IsDone = false
	toDoMessage.CreatedAt = time.Now()

	todoMessageArray.toDoMessages = append(todoMessageArray.toDoMessages, toDoMessage)

	return toDoMessage, nil
}
