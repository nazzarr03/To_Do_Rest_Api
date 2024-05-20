package services

import (
	"github.com/nazzarr03/To-Do-Rest-Api/dto"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
)

type DefaultToDoMessageService struct {
	Repo repository.ToDoMessageRepository
}

type ToDoMessageService interface {
	GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error)
	CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (*dto.ToDoMessageDto, error)
	DeleteToDoMessageByMessageID(messageID, listID, userID uint) error
}

func (toDoMessageService *DefaultToDoMessageService) GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error) {
	toDoMessages, err := toDoMessageService.Repo.GetToDoMessagesByListID(listID)
	if err != nil {
		return nil, err
	}

	return toDoMessages, nil
}

func (toDoMessageService *DefaultToDoMessageService) CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (*dto.ToDoMessageDto, error) {
	toDoMessage, err := toDoMessageService.Repo.CreateToDoMessageByListID(listID, userID, toDoMessage)
	if err != nil {
		return nil, err
	}

	toDoMessageDto := dto.ToDoMessageDto{
		MessageID: toDoMessage.MessageID,
		ListID:    toDoMessage.ListID,
		UserID:    toDoMessage.UserID,
		Content:   toDoMessage.Content,
		IsDone:    toDoMessage.IsDone,
		CreatedAt: toDoMessage.CreatedAt,
		UpdatedAt: toDoMessage.UpdatedAt,
	}

	return &toDoMessageDto, nil
}

func (toDoMessageService *DefaultToDoMessageService) DeleteToDoMessageByMessageID(messageID, listID, userID uint) error {
	err := toDoMessageService.Repo.DeleteToDoMessageByMessageID(messageID, listID, userID)
	if err != nil {
		return err
	}

	return nil
}
