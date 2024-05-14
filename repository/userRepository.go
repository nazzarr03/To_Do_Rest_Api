package repository

import "github.com/nazzarr03/To-Do-Rest-Api/models"

type UserArray struct {
	users []models.User
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
