package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) any {
	userID, exists := ctx.Get("UserID")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return nil
	}

	return userID
}
