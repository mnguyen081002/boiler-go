package routes

import "erp/internal/lib"

type UserRoutes struct {
	handler *lib.Handler
}

func NewUserRoutes(handler *lib.Handler) *UserRoutes {
	_ = handler.Group("/user")
	return &UserRoutes{
		handler: handler,
	}
}
