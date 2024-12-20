package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, "pinged reg user")
}

func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "pinged user info")
}
