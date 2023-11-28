package routes

import (
	"erp/internal/lib"
	"erp/internal/modules/auth/controllers"
)

type AuthRoutes struct {
	handler *lib.Handler
}

func NewAuthRoutes(handler *lib.Handler, controller *controllers.AuthController) *AuthRoutes {
	g := handler.Group("/auth")
	g.POST("/register", controller.Register)
	// g.POST("/login", controller.Login)
	return &AuthRoutes{
		handler: handler,
	}
}
