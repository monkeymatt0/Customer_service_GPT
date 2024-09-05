package middlewares

import (
	"customer_service_gpt/db"
	"customer_service_gpt/models"
	"customer_service_gpt/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the token is inside the DB, if it is is actually a logged person otherwise should be rejected
		var userSession models.UserSession
		if err := db.DB.Where("token = ?", token).First(&userSession).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session not found"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
