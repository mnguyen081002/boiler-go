package repository

import (
	"erp/config"
	"erp/internal/domain"
	"erp/internal/repository/gormlib"
	"erp/internal/repository/mongo"
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
	return &UnitOfWork{
		TokenRepository: mongo.NewTokenRepository(),
		UserRepository:  mongo.NewUserRepository(),
	}
}

func NewUnitOfWork(config *config.Config) *UnitOfWork {
	if config.Database.Driver == "mongo" {
		return NewUnitOfWorkMongo()
	} else {
		return NewUnitOfWorkGorm()
	}
}

var Module = fx.Options(
	//gormlib.Provides,
	fx.Provide(NewUnitOfWork),
	fx.Provide(mongo.NewMongoRepository),
)
