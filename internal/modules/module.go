package modules

import (
	"erp/internal/modules/auth"
	"erp/internal/modules/user"

	"go.uber.org/fx"
)


var Module = fx.Options(auth.Module,user.Module)
