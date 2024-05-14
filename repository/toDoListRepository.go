package repository

import (
	"sync"

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
