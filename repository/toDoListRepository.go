package repository

import (
	"time"

	"github.com/nazzarr03/To-Do-Rest-Api/config"
	"github.com/nazzarr03/To-Do-Rest-Api/helper"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoListMockRepository struct{}

type ToDoListRepository interface {
	GetTodoListsByUserID(userID uint) ([]models.ToDoList, error)
	CreateToDoList(userID uint, toDoList models.ToDoList) (models.ToDoList, error)
	DeleteToDoList(listID, userID uint) error
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

func (toDoListMockRepo *ToDoListMockRepository) CreateToDoList(userID uint, toDoList models.ToDoList) (models.ToDoList, error) {
	toDoList.ListID = helper.GetMaxListID()
	toDoList.UserID = userID
	toDoList.CompletionPercent = 0
	toDoList.CreatedAt = time.Now()
	toDoList.UpdatedAt = time.Now()

	config.ToDoLists = append(config.ToDoLists, toDoList)

	return toDoList, nil
}

func (toDoListMockRepo *ToDoListMockRepository) DeleteToDoList(listID, userID uint) error {
	var toDoList *models.ToDoList
	var user *models.User

	toDoList, err := helper.FindListByID(listID)
	if err != nil {
		return err
	}

	user, err = helper.FindUserByID(userID)
	if err != nil {
		return err
	}

	err = helper.IsAuthorized(*user, *toDoList)
	if err != nil {
		return err
	}

	toDoList.DeletedAt = time.Now()

	return nil
}
