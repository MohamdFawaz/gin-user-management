package routes

import (
	controllers2 "gin-user-management/api/v1/controllers"
	"gin-user-management/api/v1/middlewares"
	"gin-user-management/lib"
)

type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers2.UserController
	env            lib.Env
	jwtMiddleware  middlewares.JWTMiddleware
}

func (userRoutes UserRoutes) Setup() {
	userRoutes.logger.Zap.Info("setting up routes")
	auth := userRoutes.handler.Gin.Group("/api/" + userRoutes.env.APIVersion + "/").Use(userRoutes.jwtMiddleware.Handler())
	{
		auth.GET("/profile", userRoutes.userController.Profile)
		auth.POST("/image", userRoutes.userController.UploadImage)
		auth.GET("/websocket", userRoutes.userController.InitWB)
	}
}

func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers2.UserController,
	env lib.Env,
	jwtMiddleware middlewares.JWTMiddleware) UserRoutes {
	return UserRoutes{
		logger:         logger,
		handler:        handler,
		userController: userController,
		env:            env,
		jwtMiddleware: jwtMiddleware,
	}
}
