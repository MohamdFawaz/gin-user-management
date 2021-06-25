package controllers

import (
	"fmt"
	"gin-user-management/lib"
	"gin-user-management/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"path/filepath"
)

type UserController struct {
	service          services.UserService
	logger           lib.Logger
	websocketHandler lib.SocketHandler
}

func NewUserController(service services.UserService, logger lib.Logger, websocketHandler lib.SocketHandler) UserController {
	return UserController{
		service:          service,
		logger:           logger,
		websocketHandler: websocketHandler,
	}
}

func (userController UserController) Profile(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		userController.logger.Zap.Error("userId not found")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Profile Not Found",
		})
		return

	}
	userController.logger.Zap.Info(userID)
	user, err := userController.service.GetById(userID)
	if err != nil {
		userController.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (userController UserController) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("file err : %s", err.Error())})
		return
	}
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	publicFolder := filepath.Join(".", "public")
	err = os.MkdirAll(publicFolder, os.ModePerm)
	if err != nil {
		userController.logger.Zap.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := c.SaveUploadedFile(file, "public/"+newFileName); err != nil {
		userController.logger.Zap.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}



var webSocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (userController UserController) InitWB(c *gin.Context) {
	userController.websocketHandler.Setup(c)
	 c.JSON(http.StatusOK, "success")
}