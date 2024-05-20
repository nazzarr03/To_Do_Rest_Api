package app

import (
	"net/http"
	"strconv"

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
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": toDoListsDto,
	})
}

func (toDoListHandler *ToDoListHandler) CreateToDoList(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	toDoListDto, err := toDoListHandler.Service.CreateToDoList(userID.(uint))
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "To-Do list created successfully",
		"data":    toDoListDto,
	})
}

func (toDoListHandler *ToDoListHandler) DeleteToDoList(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	listIDStr := ctx.Param("listID")
	listID64, err := strconv.ParseUint(listIDStr, 10, 64)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	listID := uint(listID64)

	err = toDoListHandler.Service.DeleteToDoList(listID, userID.(uint))
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "To-Do list deleted successfully",
	})
}
