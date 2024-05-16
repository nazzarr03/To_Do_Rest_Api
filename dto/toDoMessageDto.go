package dto

import (
	"time"
)

type ToDoMessageDto struct {
	MessageID uint   `json:"message_id"`
	ListID    uint   `json:"list_id"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	IsDone    bool   `json:"is_done"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
