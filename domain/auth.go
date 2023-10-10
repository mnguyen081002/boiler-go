package domain

import (
	"context"
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

type AuthService interface {
	Register(ctx context.Context, req RegisterRequest) (user *models.User, err error)
}
