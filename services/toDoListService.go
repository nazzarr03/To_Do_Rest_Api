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
	CreateToDoList(userID uint, toDoList models.ToDoList) (*dto.ToDoListDto, error)
	DeleteToDoList(listID, userID uint) error
}

func (t *DefaultToDoListService) CreateToDoList(userID uint, toDoList models.ToDoList) (*dto.ToDoListDto, error) {
	toDoList, err := t.Repo.CreateToDoList(userID, toDoList)
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

func (t *DefaultToDoListService) DeleteToDoList(listID, userID uint) error {
	err := t.Repo.DeleteToDoList(listID, userID)
	if err != nil {
		return err
	}

	return nil
}
