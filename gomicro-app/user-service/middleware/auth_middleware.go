package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ozturkeniss/gomicro-app/user-service/auth"
)

// AuthMiddleware is a middleware that validates JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token geçerliyse, claims'i context'e ekle
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
} 