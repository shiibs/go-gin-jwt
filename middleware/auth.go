package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shiibs/go-gin-jwt/helper"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("token")

	if token == "" {
		c.JSON(401, gin.H{"error": "TOken not present."})
		c.Abort()
		return
	}

	claims, msg := helper.ValidateToken(token)

	if msg != "" {
		c.JSON(401, gin.H{"error": msg})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)

	c.Next()
}