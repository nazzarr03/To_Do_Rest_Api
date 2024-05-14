package models

import "time"

type ToDoList struct {
	ListID            uint          `json:"list_id"`
	UserID            uint          `json:"user_id"`
	CompletionPercent uint          `json:"completion_percent"`
	ToDoMessages      []ToDoMessage `json:"to_do_messages"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
