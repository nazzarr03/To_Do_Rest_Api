package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoMessageArray struct {
	toDoMessages []models.ToDoMessage
	mutex        sync.Mutex
}

type ToDoMessageRepository interface {
	CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error)
	DeleteToDoMessageByMessageID(messageID, userID uint) error
	UpdateToDoMessageByMessageID(messageID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error)
}

func (toDoMessageArray *ToDoMessageArray) CreateToDoMessageByListID(listID, userID uint, toDoMessage models.ToDoMessage) (models.ToDoMessage, error) {
	toDoMessageArray.mutex.Lock()
	defer toDoMessageArray.mutex.Unlock()

	toDoListArray := ToDoListArray{}
	userArray := UserArray{}

	for _, toDoList := range toDoListArray.toDoLists {
		if toDoList.ListID == listID {
			toDoMessage.ListID = listID
		} else {
			return models.ToDoMessage{}, errors.New("to-do list not found")
		}
	}

	for _, user := range userArray.users {
		if user.UserID == userID {
			toDoMessage.UserID = userID
		} else {
			return models.ToDoMessage{}, errors.New("user not found")
		}
	}

	toDoMessage.MessageID = uint(toDoMessageArray.toDoMessages[len(toDoMessageArray.toDoMessages)-1].MessageID) + 1
	toDoMessage.ListID = listID
	toDoMessage.IsDone = false
	toDoMessage.CreatedAt = time.Now()

	toDoMessageArray.toDoMessages = append(toDoMessageArray.toDoMessages, toDoMessage)

	return toDoMessage, nil
}

func (toDoMessageArray *ToDoMessageArray) DeleteToDoMessageByMessageID(messageID, userID uint) error {
	toDoMessageArray.mutex.Lock()
	defer toDoMessageArray.mutex.Unlock()

	toDoListArray := ToDoListArray{}
	userArray := UserArray{}

	for i, toDoMessage := range toDoMessageArray.toDoMessages {
		for _, toDoList := range toDoListArray.toDoLists {
			for _, user := range userArray.users {
				if toDoMessage.MessageID == messageID {
					if toDoList.ListID == toDoMessage.ListID {
						if user.UserID == userID {
							if user.UserType == "admin" || user.UserID == toDoMessage.UserID {
								toDoMessageArray.toDoMessages[i].DeletedAt = time.Now()

								doneCount := 0
								deleteCount := 0

								for _, toDoMessage := range toDoMessageArray.toDoMessages {
									if toDoMessage.IsDone && !toDoMessage.DeletedAt.IsZero() {
										doneCount++
									}
								}

								for _, toDoMessage := range toDoMessageArray.toDoMessages {
									if toDoMessage.DeletedAt.IsZero() {
										deleteCount++
									}
								}

								toDoListArray.toDoLists = append(toDoListArray.toDoLists, models.ToDoList{
									CompletionPercent: uint(doneCount / (len(toDoMessageArray.toDoMessages) - deleteCount) * 100),
								})

								return nil

							} else {
								return errors.New("you are not authorized to delete this message")
							}
						} else {
							return errors.New("user not found")
						}
					} else {
						return errors.New("to-do list not found")
					}
				} else {
					return errors.New("message not found")
				}
			}
		}
	}

	return nil
}
