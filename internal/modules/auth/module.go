package auth

import (
	"erp/internal/modules/auth/controllers"
	"erp/internal/modules/auth/routes"
	"erp/internal/modules/auth/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(usecases.Module,routes.Module, controllers.Module)
