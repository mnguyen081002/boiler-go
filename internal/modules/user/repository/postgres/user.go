package postgresrepository

import (
	"context"
	"erp/internal/api_errors"
	"erp/internal/domain"
	"erp/internal/infrastructure"
	"erp/internal/models"
	"erp/internal/utils"

	"github.com/pkg/errors"
)

type UserRepositoryImpl struct {
	*infrastructure.Database
}

func NewUserRepository(db *infrastructure.Database) domain.UserRepository {
	utils.MustHaveDb(db)
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, user *models.User) (res *models.User, err error) {
	err = u.DB.Create(&user).Error

	return user, err
}

func (u *UserRepositoryImpl) GetByID(ctx context.Context, id string) (res *models.User, err error) {
	err = u.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return res, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return
}

func (u *UserRepositoryImpl) IsExistEmail(ctx context.Context, email string) (res *models.User, err error) {
	err = u.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return res, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return
}
