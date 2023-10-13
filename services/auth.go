package service

import (
	"context"
	"erp/api_errors"
	config "erp/config"
	"erp/domain"
	models "erp/models"
	"erp/utils"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userService domain.UserService
	config      *config.Config
}

func NewAuthService(userService domain.UserService, config *config.Config) domain.AuthService {
	return &authService{
		userService: userService,
		config:      config,
	}
}

func (a *authService) Login(ctx context.Context, req domain.LoginInput) (result *domain.LoginResult, err error) {

	user, err := a.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New(api_errors.ErrInvalidPassword)
	}

	return &domain.LoginResult{
		User: &domain.LoginUserResult{
			Email: user.Email,
		},
		AccessToken: map[string]interface{}{
			"aud": a.config.Jwt.Aud,
			"iss": a.config.Jwt.Iss,
			"sub": user.ID,
			"jti": utils.GenerateRandomString(20),
			"exp": time.Now().Add(time.Duration(a.config.Jwt.AccessTokenExpiresIn)).Unix(),
		},
		RefreshToken: map[string]interface{}{
			"aud": a.config.Jwt.Aud,
			"iss": a.config.Jwt.Iss,
			"sub": user.ID,
			"jti": utils.GenerateRandomString(26),
			"exp": time.Now().Add(time.Duration(a.config.Jwt.RefreshTokenExpiresIn)).Unix(),
		},
	}, nil
}

func (a *authService) Register(ctx context.Context, req domain.RegisterRequest) (user *models.User, err error) {
	user, err = a.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		if err.Error() != api_errors.ErrUserNotFound {
			return nil, err
		}
	}
	if user != nil {
		return nil, errors.New(api_errors.ErrEmailAlreadyExist)
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
		Email:    req.Email,
		Password: req.Password,
	})

	return user, err
}
