package routes

import "go.uber.org/fx"

var Module = fx.Options(fx.Invoke(NewAuthRoutes))