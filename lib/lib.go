package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewSocketHandler),
	fx.Provide(NewCronjob),
)
