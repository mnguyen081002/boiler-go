package usecases

import (
	"context"
	config "erp/config"
	"erp/internal/domain"
	models "erp/internal/models"
)

type (
	UserServiceImpl struct {
		userRepo domain.UserRepository
		config   *config.Config
	}
)

// Create implements UserService.
func (u *UserServiceImpl) Create(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := u.userRepo.Create(ctx, user)
	return user, err
}

func NewUserService(userRepo domain.UserRepository, config *config.Config) domain.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id string) (user *models.User, err error) {
	user, err = s.userRepo.GetByID(ctx, id)
	return
}
