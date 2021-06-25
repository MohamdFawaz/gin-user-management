package routes

import (
	controllers2 "gin-user-management/api/v1/controllers"
	"gin-user-management/lib"
)

type AuthRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	authController controllers2.AuthController
	env            lib.Env
}

func (authRoutes AuthRoutes) Setup() {
	authRoutes.logger.Zap.Info("setting up routes")
	auth := authRoutes.handler.Gin.Group("/api/" + authRoutes.env.APIVersion + "/auth")
	{
		auth.POST("/login", authRoutes.authController.SignIn)
		auth.POST("/register", authRoutes.authController.Register)
	}
}

func NewAuthRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	authController controllers2.AuthController,
	env lib.Env) AuthRoutes {
	return AuthRoutes{
		logger:         logger,
		handler:        handler,
		authController: authController,
		env: env,
	}
}
