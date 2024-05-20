package services

import (
	"github.com/nazzarr03/To-Do-Rest-Api/dto"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
)

type DefaultToDoListService struct {
	Repo repository.ToDoListRepository
}

type ToDoListService interface {
	GetTodoListsByUserID(userID uint) ([]models.ToDoList, error)
	CreateToDoList(userID uint) (*dto.ToDoListDto, error)
	DeleteToDoList(listID, userID uint) error
}

func (toDoListService *DefaultToDoListService) GetTodoListsByUserID(userID uint) ([]models.ToDoList, error) {
	myTodoLists, err := toDoListService.Repo.GetTodoListsByUserID(userID)

	if err != nil {
		return nil, err
	}
	return myTodoLists, nil
}

func (toDoListService *DefaultToDoListService) CreateToDoList(userID uint) (*dto.ToDoListDto, error) {
	var toDoList models.ToDoList

	toDoList, err := toDoListService.Repo.CreateToDoList(userID, toDoList)
	if err != nil {
		return nil, err
	}

	toDoListDto := dto.ToDoListDto{
		ListID:            toDoList.ListID,
		UserID:            toDoList.UserID,
		CompletionPercent: toDoList.CompletionPercent,
		ToDoMessages:      toDoList.ToDoMessages,
		CreatedAt:         toDoList.CreatedAt,
		UpdatedAt:         toDoList.UpdatedAt,
	}

	return &toDoListDto, nil
}

func (toDoListService *DefaultToDoListService) DeleteToDoList(listID, userID uint) error {
	err := toDoListService.Repo.DeleteToDoList(listID, userID)
	if err != nil {
		return err
	}

	return nil
}
