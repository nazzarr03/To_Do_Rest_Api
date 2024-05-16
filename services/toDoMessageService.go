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
	CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (*dto.ToDoMessageDto, error)
	DeleteToDoMessageByMessageID(messageID, userID uint) error
	UpdateToDoMessageByMessageID(messageID, userID uint, toDoMessage models.ToDoMessage) (*dto.ToDoMessageDto, error)
}

func (t *DefaultToDoMessageService) CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (*dto.ToDoMessageDto, error) {
	toDoMessage, err := t.Repo.CreateToDoMessageByListID(listID, userID, toDoMessage)
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

func (t *DefaultToDoMessageService) DeleteToDoMessageByMessageID(messageID, userID uint) error {
	err := t.Repo.DeleteToDoMessageByMessageID(messageID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (t *DefaultToDoMessageService) UpdateToDoMessageByMessageID(messageID, userID uint, toDoMessage models.ToDoMessage) (*dto.ToDoMessageDto, error) {
	toDoMessage, err := t.Repo.UpdateToDoMessageByMessageID(messageID, userID, toDoMessage)
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
