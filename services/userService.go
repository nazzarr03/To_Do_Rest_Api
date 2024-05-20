package services

import (
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	Login(username, password string) (string, error)
}

func (userService *DefaultUserService) Login(username, password string) (string, error) {
	token, err := userService.Repo.Login(username, password)
	if err != nil {
		return "", err
	}

	return token, nil
}
