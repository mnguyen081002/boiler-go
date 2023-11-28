package user

import (
	"erp/internal/modules/user/controllers"
	"erp/internal/modules/user/repository"
	"erp/internal/modules/user/routes"
	"erp/internal/modules/user/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(repository.Module, usecases.Module, routes.Module, controllers.Module)
