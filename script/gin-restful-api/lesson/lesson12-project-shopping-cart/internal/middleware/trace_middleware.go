package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"trungem.com/shopping-cart/pkg/logger"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		contextValue := context.WithValue(ctx.Request.Context(), logger.TraceIdKey, traceID)

		ctx.Request = ctx.Request.WithContext(contextValue)

		ctx.Writer.Header().Set("X-Trace-ID", traceID)

		ctx.Set(string(logger.TraceIdKey), traceID)

		ctx.Next()
	}
}
