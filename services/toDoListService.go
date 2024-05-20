package services

import (
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
)

type DefaultToDoListService struct {
	Repo repository.ToDoListRepository
}

type ToDoListService interface {
	GetTodoListsByUserID(userID uint) ([]models.ToDoList, error)
}

func (toDoListService *DefaultToDoListService) GetTodoListsByUserID(userID uint) ([]models.ToDoList, error) {
	myTodoLists, err := toDoListService.Repo.GetTodoListsByUserID(userID)

	if err != nil {
		return nil, err
	}
	return myTodoLists, nil
}
