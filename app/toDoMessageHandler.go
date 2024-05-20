package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/helper"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/services"
)

type ToDoMessageHandler struct {
	Service services.ToDoMessageService
}

func (toDoMessageHandler *ToDoMessageHandler) GetToDoMessagesByListID(ctx *gin.Context) {
	listID, err := helper.ParseUintParam(ctx, "listID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	toDoMessages, err := toDoMessageHandler.Service.GetToDoMessagesByListID(listID)
	if err != nil {
		helper.HandleError(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": toDoMessages,
	})
}

func (toDoMessageHandler *ToDoMessageHandler) CreateToDoMessageByListID(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	listID, err := helper.ParseUintParam(ctx, "listID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	var toDoMessage models.ToDoMessage
	if err := ctx.ShouldBindJSON(&toDoMessage); err != nil {
		helper.HandleError(ctx, err)
	}

	toDoMessageDto, err := toDoMessageHandler.Service.CreateToDoMessageByListID(listID, userID.(uint), toDoMessage)
	if err != nil {
		helper.HandleError(ctx, err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "To-Do message created successfully",
		"data":    toDoMessageDto,
	})
}

func (toDoMessageHandler *ToDoMessageHandler) DeleteToDoMessageByMessageID(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	messageID, err := helper.ParseUintParam(ctx, "messageID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	listID, err := helper.ParseUintParam(ctx, "listID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	err = toDoMessageHandler.Service.DeleteToDoMessageByMessageID(messageID, listID, userID.(uint))
	if err != nil {
		helper.HandleError(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "To-Do message deleted successfully",
	})
}

func (toDoMessageHandler *ToDoMessageHandler) UpdateToDoMessageByMessageID(ctx *gin.Context) {
	userID := helper.GetUserID(ctx)

	var toDoMessage models.ToDoMessage
	if err := ctx.ShouldBindJSON(&toDoMessage); err != nil {
		helper.HandleError(ctx, err)
	}

	messageID, err := helper.ParseUintParam(ctx, "messageID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	listID, err := helper.ParseUintParam(ctx, "listID")
	if err != nil {
		helper.HandleError(ctx, err)
	}

	toDoMessageDto, err := toDoMessageHandler.Service.UpdateToDoMessageByMessageID(messageID, listID, userID.(uint), toDoMessage)
	if err != nil {
		helper.HandleError(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "To-Do message updated successfully",
		"data":    toDoMessageDto,
	})
}
