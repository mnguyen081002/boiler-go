package controllers

import (
	"erp/internal/api"
	"erp/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService domain.AuthService
	logger      *zap.Logger
}

func NewAuthController(authService domain.AuthService, logger *zap.Logger) *AuthController {
	controller := &AuthController{
		AuthService: authService,
		logger:      logger,
	}
	return controller
}

func (b *AuthController) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.ResponseError(c, err)
		return
	}

	_, err := b.AuthService.Register(c.Request.Context(), req)

	if err != nil {
		api.ResponseError(c, err)
		return
	}
	api.Response(c, http.StatusOK, "success", nil)
}
