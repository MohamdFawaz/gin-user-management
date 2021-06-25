package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	authRoutes AuthRoutes,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		authRoutes,
		userRoutes,
	}
}

func (routes Routes) Setup() {
	for _, route := range routes {
		route.Setup()
	}
}
