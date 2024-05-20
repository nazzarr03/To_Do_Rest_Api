package services

import (
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
)

type DefaultToDoMessageService struct {
	Repo repository.ToDoMessageRepository
}

type ToDoMessageService interface {
	GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error)
}

func (toDoMessageService *DefaultToDoMessageService) GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error) {
	toDoMessages, err := toDoMessageService.Repo.GetToDoMessagesByListID(listID)
	if err != nil {
		return nil, err
	}

	return toDoMessages, nil
}
