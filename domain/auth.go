package domain

import (
	"context"
	"erp/constants"
	"erp/models"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	jwt.StandardClaims
	// RoleID    string              `json:"role_id"`
	TokenType string `json:"token_type"`
}

type RegisterRequest struct {
	Email       string `json:"email" binding:"required" validate:"email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	FirstName   string `json:"first_name" binding:"required" validate:"min=1,max=50"`
	LastName    string `json:"last_name" binding:"required" validate:"min=1,max=50"`
	RequestFrom string `json:"request_from" binding:"required" enums:"erp/,web,app"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type LoginResult struct {
	User         *models.User           `json:"user"`
	AccessToken  map[string]interface{} `json:"access_token"`
	RefreshToken map[string]interface{} `json:"refresh_token"`
}

type AuthService interface {
	Register(ctx context.Context, req RegisterRequest) (user *models.User, err error)
	Login(ctx context.Context, req LoginInput) (result *LoginResult, err error)
}

type JwtService interface {
	GenerateToken(userID string, kid string, tokenType constants.TokenType, expiresIn int64) (string, error)
	ValidateToken(token string, tokenType constants.TokenType) (*string, error)
	GenerateAuthTokens(userID string) (string, string, error)
}
