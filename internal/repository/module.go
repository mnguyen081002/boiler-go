package repository

import (
	"erp/internal/domain"
	"erp/internal/repository/gormlib"
	"go.uber.org/fx"
)

type UnitOfWork struct {
	TokenRepository domain.TokenRepository
	UserRepository  domain.UserRepository
}

func NewUnitOfWorkGorm() *UnitOfWork {
	return &UnitOfWork{
		TokenRepository: gormlib.NewTokenRepository(),
		UserRepository:  gormlib.NewUserRepository(),
	}
}

func NewUnitOfWorkMongo() *UnitOfWork {
	return &UnitOfWork{}
}

var Module = fx.Options(
	//gormlib.Provides,
	fx.Provide(NewUnitOfWorkGorm),
	fx.Provide(gormlib.NewGormRepository),
)
