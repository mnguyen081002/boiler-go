package bootstrap

import (
	"erp/config"
	"erp/internal/api/middlewares"
	"erp/internal/infrastructure"
	"erp/internal/lib"
	"erp/internal/modules"
	"erp/internal/utils"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func inject() fx.Option {
	return fx.Options(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		//fx.NopLogger,
		fx.Provide(
			config.NewConfig,
			utils.NewTimeoutContext,
		),
		lib.Module,
		middlewares.Module,
		infrastructure.Module,
		modules.Module,
	)
}

func Run() {
	fx.New(inject()).Run()
}
