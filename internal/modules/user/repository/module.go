package repository

import (
	postgresrepository "erp/internal/modules/user/repository/postgres"

	"go.uber.org/fx"
)

var Module = fx.Options(postgresrepository.Module)
