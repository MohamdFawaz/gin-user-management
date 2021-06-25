package controllers

import (
	"gin-user-management/lib"
	"gin-user-management/models"
	"gin-user-management/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	service services.AuthService
	logger  lib.Logger
}

// NewAuthController creates new auth controller
func NewAuthController(service services.AuthService, logger lib.Logger) AuthController {
	return AuthController{
		service: service,
		logger:  logger,
	}
}

func (authController AuthController) SignIn(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		authController.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, user, err := authController.service.Login(user.Username)
	if err != nil {
		authController.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (authController AuthController) Register(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		authController.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := authController.service.CreateNewUser(user)
	if err != nil {
		authController.logger.Zap.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := authController.service.CreateToken(user)
	if err != nil {
		authController.logger.Zap.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
