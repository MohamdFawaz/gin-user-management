package bootstrap

import (
	"context"
	"gin-user-management/api/v1/controllers"
	"gin-user-management/api/v1/middlewares"
	"gin-user-management/api/v1/routes"
	"gin-user-management/lib"
	"gin-user-management/repository"
	"gin-user-management/services"
	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	repository.Module,
	middlewares.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	logger lib.Logger,
	middlewares middlewares.Middlewares,
	cronjob lib.Cronjob,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")

			conn.SetMaxOpenConns(10)

			go func() {
				routes.Setup()
				middlewares.Setup()
				cronjob.SetupJobs()
				logger.Zap.Error(handler.Gin.Run(":" + env.ServerPort))
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
