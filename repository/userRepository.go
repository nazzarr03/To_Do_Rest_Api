package repository

import (
	"errors"

	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/utils"
)

type UserArray struct {
	users []models.User
}

type UserRepository interface {
	Login(user models.User) (string, error)
}

func AddUser() {
	userArray := UserArray{}

	userArray.users = append(userArray.users, models.User{
		UserID:   1,
		Username: "username1",
		Password: "password1",
		UserType: "default",
	}, models.User{
		UserID:   2,
		Username: "username2",
		Password: "password2",
		UserType: "admin",
	})
}

func (userArray *UserArray) Login(user models.User) (string, error) {
	for _, u := range userArray.users {
		if u.Username == user.Username && u.Password == user.Password {
			token, err := utils.GenerateToken(u.UserID)
			if err != nil {
				return "", err
			}

			return token, nil
		}
	}

	return "", errors.New("invalid username or password")
}
