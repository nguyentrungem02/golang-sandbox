package app

import (
	v1handler "trungem.com/shopping-cart/internal/handler/v1"
	"trungem.com/shopping-cart/internal/repository"
	"trungem.com/shopping-cart/internal/routes"
	v1routes "trungem.com/shopping-cart/internal/routes/v1"
	v1service "trungem.com/shopping-cart/internal/service/v1"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(ctx *ModuleContext) *UserModule {
	// Initialize repository
	userRepo := repository.NewSqlUserRepository(ctx.DB)

	// Initialize service
	userService := v1service.NewUserService(userRepo, ctx.Redis)

	// Initialize handler
	userHandler := v1handler.NewUserHandler(userService)

	// Initialize routes
	userRoutes := v1routes.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}
