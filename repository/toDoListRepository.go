package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoListArray struct {
	toDoLists []models.ToDoList
	mutex     sync.Mutex
}

type ToDoListRepository interface {
	CreateToDoList(userID uint, toDoList models.ToDoList) (models.ToDoList, error)
	DeleteToDoList(listID, userID uint) error
}

func (toDoListArray *ToDoListArray) CreateToDoList(userID uint, toDoList models.ToDoList) (models.ToDoList, error) {
	toDoListArray.mutex.Lock()
	defer toDoListArray.mutex.Unlock()

	UserArray := UserArray{}

	for _, user := range UserArray.users {
		if user.UserID == userID {
			toDoList.UserID = userID
		} else {
			return models.ToDoList{}, errors.New("user not found")
		}
	}

	toDoList.ListID = uint(toDoListArray.toDoLists[len(toDoListArray.toDoLists)-1].ListID) + 1
	toDoList.CompletionPercent = 0
	toDoList.CreatedAt = time.Now()

	toDoListArray.toDoLists = append(toDoListArray.toDoLists, toDoList)

	return toDoList, nil
}
