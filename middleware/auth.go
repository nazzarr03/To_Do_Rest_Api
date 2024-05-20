package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/To-Do-Rest-Api/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing please enter a token."})
			c.Abort()
			return
		}
		authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")[1]

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(authHeader)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("UserID", claims.UserId)
		c.Next()
	}
}
