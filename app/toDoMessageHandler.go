package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/services"
)

type ToDoMessageHandler struct {
	Service services.ToDoMessageService
}

func (toDoMessageHandler *ToDoMessageHandler) GetToDoMessagesByListID(ctx *gin.Context) {
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

	toDoMessages, err := toDoMessageHandler.Service.GetToDoMessagesByListID(listID)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": toDoMessages,
	})
}
