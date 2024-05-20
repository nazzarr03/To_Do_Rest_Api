package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/helper"
	"github.com/nazzarr03/To-Do-Rest-Api/services"
)

type ToDoListHandler struct {
	Service services.ToDoListService
}

func (toDoListHandler *ToDoListHandler) GetTodoListsByUserID(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	toDoListsDto, err := toDoListHandler.Service.GetTodoListsByUserID(userID.(uint))

	helper.HandleError(ctx, err)

	ctx.JSON(http.StatusOK, gin.H{
		"data": toDoListsDto,
	})
}

func (toDoListHandler *ToDoListHandler) CreateToDoList(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	toDoListDto, err := toDoListHandler.Service.CreateToDoList(userID.(uint))

	helper.HandleError(ctx, err)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "To-Do list created successfully",
		"data":    toDoListDto,
	})
}

func (toDoListHandler *ToDoListHandler) DeleteToDoList(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	listID, err := helper.ParseUintParam(ctx, "listID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	err = toDoListHandler.Service.DeleteToDoList(listID, userID.(uint))

	helper.HandleError(ctx, err)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "To-Do list deleted successfully",
	})
}
