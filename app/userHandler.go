package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/helper"
	"github.com/nazzarr03/To-Do-Rest-Api/models"
	"github.com/nazzarr03/To-Do-Rest-Api/services"
)

type UserHandler struct {
	Service services.UserService
}

func (userHandler *UserHandler) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helper.HandleError(ctx, err)
	}

	token, err := userHandler.Service.Login(user.Username, user.Password)
	if err != nil {
		helper.HandleError(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   token,
	})
}
