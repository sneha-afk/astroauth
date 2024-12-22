package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sneha-afk/astroauth/models"
	"github.com/sneha-afk/astroauth/store"
)

func RegisterUser(c *gin.Context) {
	log.Print("pinged reg user")
	var regUser models.UserInternal
	if err := c.BindJSON(&regUser); err != nil {
		log.Printf("RegisterUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Generate and set random UUID
	regUser.ID = uuid.New().String()

	if err := store.CreateUser(regUser); err != nil {
		errstr := fmt.Sprintf("RegisterUser: %v", err)
		log.Print(errstr)
		c.JSON(http.StatusBadRequest, gin.H{"error": errstr})
		return
	}

	// Return back UUID
	c.JSON(http.StatusOK, gin.H{"message": "registered user", "id": regUser.ID})
}

func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "pinged user info")
}
