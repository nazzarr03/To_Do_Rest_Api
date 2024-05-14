package models

type UserType string

const (
	Default UserType = "default"
	Admin   UserType = "admin"
)

type User struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
