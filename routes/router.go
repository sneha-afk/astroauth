package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sneha-afk/astroauth/controllers"
	"github.com/sneha-afk/astroauth/utils"
)

func RegisterRoutes(r *gin.Engine) {
	// Handle panics by sending 500 (internal server err) back
	r.Use(gin.Recovery())

	// For now, only take in JSON
	r.Use(utils.ContentTypeIsJson())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/register", controllers.RegisterUser)
		v1.POST("/login", controllers.LoginUser)

		// Protected routes
		authorized := v1.Group("/")
		{
			authorized.Use(utils.AuthVerification())

			authorized.GET("/user/:username", controllers.GetUserInfo)
		}
	}

}
