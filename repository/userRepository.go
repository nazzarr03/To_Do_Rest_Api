package repository

import (
	"errors"

	"github.com/nazzarr03/To-Do-Rest-Api/config"
	"github.com/nazzarr03/To-Do-Rest-Api/utils"
)

type UserMockRepository struct{}

type UserRepository interface {
	Login(username, password string) (string, error)
}

func (userMockRepo *UserMockRepository) Login(username, password string) (string, error) {
	for _, user := range config.Users {
		if user.Username == username && user.Password == password {
			token, err := utils.GenerateToken(user.UserID)
			if err != nil {
				return "", err
			}

			return token, nil
		}
	}

	return "", errors.New("invalid username or password")
}
