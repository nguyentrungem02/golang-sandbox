package app

import (
	v1handler "trungem.com/shopping-cart/internal/handler/v1"
	"trungem.com/shopping-cart/internal/repository"
	"trungem.com/shopping-cart/internal/routes"
	v1routes "trungem.com/shopping-cart/internal/routes/v1"
	v1service "trungem.com/shopping-cart/internal/service/v1"
	"trungem.com/shopping-cart/pkg/auth"
	"trungem.com/shopping-cart/pkg/cache"
	"trungem.com/shopping-cart/pkg/mail"
	"trungem.com/shopping-cart/pkg/rabbitmq"
)

type AuthModule struct {
	routes routes.Route
}

func NewAuthModule(ctx *ModuleContext, tokenService auth.TokenService, cacheService cache.RedisCacheService, mailService mail.EmailProviderService, rabbitmqService rabbitmq.RabbitMQService) *AuthModule {
	// Initialize repository
	userRepo := repository.NewSqlUserRepository(ctx.DB)

	// Initialize service
	authService := v1service.NewAuthService(userRepo, tokenService, cacheService, mailService, rabbitmqService)

	// Initialize handler
	authHandler := v1handler.NewAuthHandler(authService)

	// Initialize routes
	authRoutes := v1routes.NewAuthRoutes(authHandler)

	return &AuthModule{
		routes: authRoutes,
	}
}

func (m *AuthModule) Routes() routes.Route {
	return m.routes
}
