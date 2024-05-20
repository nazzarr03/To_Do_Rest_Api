package config

import "github.com/nazzarr03/To-Do-Rest-Api/models"

var Users = []models.User{
	{UserID: 1, Username: "nazzarr03", Password: "password", UserType: "admin"},
	{UserID: 2, Username: "nazzarr04", Password: "password", UserType: "default"},
}

var ToDoLists = []models.ToDoList{}

var ToDoMessages = []models.ToDoMessage{}
