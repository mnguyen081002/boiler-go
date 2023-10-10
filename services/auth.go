package service

import (
	"context"
	config "erp/config"
	constants "erp/constants"
	"erp/domain"
	models "erp/models"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userService domain.UserService
	config      *config.Config
}

func NewAuthService(userService domain.UserService, config *config.Config) domain.AuthService {
	return &AuthServiceImpl{
		userService: userService,
		config:      config,
	}
}

func (a *AuthServiceImpl) Register(ctx context.Context, req domain.RegisterRequest) (user *models.User, err error) {
	role := constants.RoleCustomer

	if req.RequestFrom != string(constants.Web) {
		role = constants.RoleSeller
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encrypt password")
	}

	req.Password = string(encryptedPassword)
	user, err = a.userService.Create(ctx, &models.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      role,
	})

	return user, err
}
