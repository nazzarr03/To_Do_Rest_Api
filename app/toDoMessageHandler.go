package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/helper"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
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

func (toDoMessageHandler *ToDoMessageHandler) CreateToDoMessageByListID(ctx *gin.Context) {
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

	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	var toDoMessage models.ToDoMessage
	if err := ctx.ShouldBindJSON(&toDoMessage); err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	toDoMessageDto, err := toDoMessageHandler.Service.CreateToDoMessageByListID(listID, userID.(uint), toDoMessage)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "To-Do message created successfully",
		"data":    toDoMessageDto,
	})
}

func (toDoMessageHandler *ToDoMessageHandler) DeleteToDoMessageByMessageID(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	messageIDStr := ctx.Param("messageID")
	messageID64, err := strconv.ParseUint(messageIDStr, 10, 64)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	messageID := uint(messageID64)

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

	err = toDoMessageHandler.Service.DeleteToDoMessageByMessageID(messageID, listID, userID.(uint))
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "To-Do message deleted successfully",
	})
}
