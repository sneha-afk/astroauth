package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sneha-afk/astroauth/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", controllers.RegisterUser)
	r.GET("/user/:id", controllers.GetUserInfo)
}
