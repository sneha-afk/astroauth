package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sneha-afk/astroauth/models"
)

func RegisterUser(c *gin.Context) {
	log.Print("pinged reg user")
	var regUser models.UserInternal
	if err := c.BindJSON(&regUser); err != nil {
		log.Printf("RegisterUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered user", "username": regUser.Username})
}

func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "pinged user info")
}
