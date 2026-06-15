package app

import (
	"trungem.com/user-manager/internal/handler"
	"trungem.com/user-manager/internal/repository"
	"trungem.com/user-manager/internal/routes"
	"trungem.com/user-manager/internal/service"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	// Initialize repository
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)

	// Initialize routes
	userRoutes := routes.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}
