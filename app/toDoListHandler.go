package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/services"
)

type ToDoListHandler struct {
	Service services.ToDoListService
}

func (toDoListHandler *ToDoListHandler) GetTodoListsByUserID(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "user not authorized",
		})
		return
	}

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
