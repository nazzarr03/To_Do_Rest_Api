package repository

import (
	"time"

	"github.com/nazzarr03/To-Do-Rest-Api/config"
	"github.com/nazzarr03/To-Do-Rest-Api/helper"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoMessageMockRepository struct{}

type ToDoMessageRepository interface {
	GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error)
	CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error)
	DeleteToDoMessageByMessageID(messageID, listID, userID uint) error
}

func (toDoMessageMockRepo *ToDoMessageMockRepository) GetToDoMessagesByListID(listID uint) ([]models.ToDoMessage, error) {
	var toDoMessages []models.ToDoMessage

	for _, toDoMessage := range config.ToDoMessages {
		if toDoMessage.DeletedAt.IsZero() && toDoMessage.ListID == listID {
			toDoMessages = append(toDoMessages, toDoMessage)
		}
	}

	return toDoMessages, nil
}

func (toDoMessageMockRepo *ToDoMessageMockRepository) CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error) {
	var toDoList *models.ToDoList
	var user *models.User

	toDoList, err := helper.FindListByID(listID)
	if err != nil {
		return models.ToDoMessage{}, err
	}

	user, err = helper.FindUserByID(userID)
	if err != nil {
		return models.ToDoMessage{}, err
	}

	err = helper.IsAuthorized(*user, *toDoList)
	if err != nil {
		return models.ToDoMessage{}, err
	}

	toDoList.CompletionPercent = helper.CalculateCompletionPercent(*toDoList)

	toDoMessage.MessageID = helper.GetMaxMessageID()
	toDoMessage.UserID = userID
	toDoMessage.ListID = listID
	toDoMessage.Content = toDoMessage.Content
	toDoMessage.IsDone = false
	toDoMessage.CreatedAt = time.Now()
	toDoMessage.UpdatedAt = time.Now()

	config.ToDoMessages = append(config.ToDoMessages, toDoMessage)

	toDoList.ToDoMessages = append(toDoList.ToDoMessages, toDoMessage)

	return toDoMessage, nil
}

func (toDoMessageMockRepo *ToDoMessageMockRepository) DeleteToDoMessageByMessageID(messageID, listID, userID uint) error {
	var toDoList *models.ToDoList
	var message *models.ToDoMessage
	var user *models.User

	toDoList, err := helper.FindListByID(listID)
	if err != nil {
		return err
	}

	message, err = helper.FindMessageByID(messageID)
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

	message.DeletedAt = time.Now()

	toDoList.CompletionPercent = helper.CalculateCompletionPercent(*toDoList)

	return nil
}
