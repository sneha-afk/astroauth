package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sneha-afk/astroauth/models"
	"github.com/sneha-afk/astroauth/store"
	"github.com/sneha-afk/astroauth/utils"
)

func RegisterUser(c *gin.Context) {
	var regUser models.UserInternal
	if err := c.BindJSON(&regUser); err != nil {
		log.Printf("RegisterUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Enforce password length, though client probably should take care of it
	if len(regUser.Password) < 8 || 72 < len(regUser.Password) {
		log.Printf("RegisterUser(): password length not suitable")
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be between 8 and 72 characters"})
		return
	}

	// Hash the password
	hashed, err := utils.HashPassword(regUser.Password)
	if err != nil {
		log.Printf("RegisterUser(): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	regUser.Password = string(hashed)

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
	username, _ := c.Get("username")
	username, ok := username.(string)
	if !ok {
		username = "<couldn't parse>"
	}
	c.JSON(http.StatusOK, fmt.Sprintf("pinged user info, hi %v", username))
}

func LoginUser(c *gin.Context) {
	var loginAttempt models.UserInternal
	if err := c.BindJSON(&loginAttempt); err != nil {
		log.Printf("LoginUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Check credentials with the DB
	success, err := store.CheckUserCredentials(loginAttempt)
	if !success || err != nil {
		details := err.Error()
		if err == nil {
			details = "details did not match"
		}
		log.Printf("LoginUser: %v", details)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not login", "details": details})
		return
	}

	// Give them a JWT
	userJWT := jwt.NewWithClaims(utils.SigningMethod,
		jwt.MapClaims{
			"iss": "astroauth-server",
			"sub": loginAttempt.Username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	signedToken, err := userJWT.SignedString(utils.PrivateKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not sign token", "details": err.Error()})
		log.Printf("LoginUser: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": signedToken})
}
