package dto

import (
	"time"

	"github.com/nazzarr03/To-Do-Rest-Api/models"
)

type ToDoListDto struct {
	ListID            uint                 `json:"list_id"`
	UserID            uint                 `json:"user_id"`
	CompletionPercent uint                 `json:"completion_percent"`
	ToDoMessages      []models.ToDoMessage `json:"to_do_messages"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
