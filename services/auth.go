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

type authService struct {
	userService domain.UserService
	jwtService  domain.JwtService
	config      *config.Config
}

func NewAuthService(userService domain.UserService, jwtService domain.JwtService, config *config.Config) domain.AuthService {
	return &authService{
		userService: userService,
		jwtService:  jwtService,
		config:      config,
	}
}

func (a *authService) Login(ctx context.Context, req domain.LoginInput) (result *domain.LoginResult, err error) {
	return &domain.LoginResult{
		User: &models.User{},
		AccessToken: map[string]interface{}{
			"aud": "http://api.example.com",
			"iss": "https://krakend.io",
			"sub": "1234567890qwertyuio",
			"jti": "mnb23vcsrt756yuiomnbvcx98ertyuiop",
			"exp": 1735689600,
		},
		RefreshToken: map[string]interface{}{
			"aud": "http://api.example.com",
			"iss": "https://krakend.io",
			"sub": "1234567890qwertyuio",
			"jti": "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
			"exp": 1735689600,
		},
	}, nil
}

func (a *authService) Register(ctx context.Context, req domain.RegisterRequest) (user *models.User, err error) {
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
