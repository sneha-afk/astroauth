package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContentTypeIsJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
			c.Abort()
			return
		}

		// If passed check, continue to next handler
		c.Next()
	}
}
