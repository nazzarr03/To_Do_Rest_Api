package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseUintParam(ctx *gin.Context, param string) (uint, error) {
	paramStr := ctx.Param(param)
	param64, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(param64), nil
}

func GetUserID(ctx *gin.Context) any {
	userID, exists := ctx.Get("UserID")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return nil
	}

	return userID
}

func HandleError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
