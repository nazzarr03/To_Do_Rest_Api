package services

import (
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	Login(user models.User) (string, error)
}

func (u *DefaultUserService) Login(user models.User) (string, error) {
	token, err := u.Repo.Login(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
