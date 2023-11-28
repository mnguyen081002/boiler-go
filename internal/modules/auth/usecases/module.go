package usecases

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthService,
)
