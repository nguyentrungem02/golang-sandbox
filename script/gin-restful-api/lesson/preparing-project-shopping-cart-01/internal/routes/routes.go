package routes

import (
	"github.com/gin-gonic/gin"
	"trungem.com/shopping-cart/internal/middleware"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	r.Use(
		middleware.LoggerMiddleware(),
		middleware.ApiKeyMiddleware(),
		middleware.AuthMiddleware(),
		middleware.RateLimiterMiddleware(),
	)

	v1api := r.Group("/api/v1")
	for _, route := range routes {
		route.Register(v1api)
	}
}
