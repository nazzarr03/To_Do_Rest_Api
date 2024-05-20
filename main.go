package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/app"
	"github.com/nazzarr03/To-Do-Rest-Api/middleware"
	"github.com/nazzarr03/To-Do-Rest-Api/repository"
	"github.com/nazzarr03/To-Do-Rest-Api/services"
)

func main() {
	appRouter := gin.Default()

	userRepository := &repository.UserMockRepository{}
	userService := &services.DefaultUserService{Repo: userRepository}
	userHandler := &app.UserHandler{Service: userService}

	toDoListRepository := &repository.ToDoListMockRepository{}
	toDoListService := &services.DefaultToDoListService{Repo: toDoListRepository}
	toDoListHandler := &app.ToDoListHandler{Service: toDoListService}

	toDoMessageRepository := &repository.ToDoMessageMockRepository{}
	toDoMessageService := &services.DefaultToDoMessageService{Repo: toDoMessageRepository}
	toDoMessageHandler := &app.ToDoMessageHandler{Service: toDoMessageService}

	appRouter.POST("/login", userHandler.Login)

	appRouter.Use(middleware.Authentication())

	appRouter.GET("/todolist", toDoListHandler.GetTodoListsByUserID)
	appRouter.POST("/todolist", toDoListHandler.CreateToDoList)
	appRouter.DELETE("/todolist/:listID", toDoListHandler.DeleteToDoList)

	appRouter.GET("/todomessage/:listID", toDoMessageHandler.GetToDoMessagesByListID)
	appRouter.POST("/todomessage/:listID", toDoMessageHandler.CreateToDoMessageByListID)
	appRouter.DELETE("/todomessage/:listID/:messageID", toDoMessageHandler.DeleteToDoMessageByMessageID)
	appRouter.PUT("/todomessage/:listID/:messageID", toDoMessageHandler.UpdateToDoMessageByMessageID)

	appRouter.Run(":8080")
}
