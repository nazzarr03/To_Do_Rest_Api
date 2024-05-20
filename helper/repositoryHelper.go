package helper

import (
	"errors"

	"github.com/nazzarr03/To-Do-Rest-Api/config"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

func GetMaxListID() uint {
	maxListID := uint(0)
	for _, item := range config.ToDoLists {
		if item.ListID > maxListID {
			maxListID = item.ListID
		}
	}
	return maxListID + 1
}

func FindListByID(listID uint) (*models.ToDoList, error) {
	for i, list := range config.ToDoLists {
		if list.ListID == listID {
			return &config.ToDoLists[i], nil
		}
	}
	return nil, errors.New("to-do list not found")
}

func FindUserByID(userID uint) (*models.User, error) {
	for i, u := range config.Users {
		if u.UserID == userID {
			return &config.Users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func FindMessageByID(messageID uint) (*models.ToDoMessage, error) {
	for i, message := range config.ToDoMessages {
		if message.MessageID == messageID {
			return &config.ToDoMessages[i], nil
		}
	}
	return nil, errors.New("to-do message not found")
}

func IsAuthorized(user models.User, list models.ToDoList) error {
	if user.UserType != "admin" && user.UserID != list.UserID {
		return errors.New("you are not authorized to delete this to-do list")
	}

	return nil
}

func GetMaxMessageID() uint {
	maxMessageID := uint(0)
	for _, item := range config.ToDoMessages {
		if item.MessageID > maxMessageID {
			maxMessageID = item.MessageID
		}
	}
	return maxMessageID + 1
}

func CalculateCompletionPercent(list models.ToDoList) uint {
	doneCount := 0
	deleteCount := 0
	messageCount := len(list.ToDoMessages)

	for _, toDoMessage := range list.ToDoMessages {
		if toDoMessage.DeletedAt.IsZero() {
			if toDoMessage.IsDone {
				doneCount++
			} else {
				deleteCount++
			}
		}
	}
	if doneCount == 0 || messageCount-deleteCount == 0 {
		return 0
	}

	return uint((uint(doneCount) / uint(messageCount-deleteCount)) * 100)
}
