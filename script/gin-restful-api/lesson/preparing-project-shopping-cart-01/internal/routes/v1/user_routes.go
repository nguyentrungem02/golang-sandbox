package v1routes

import (
	"github.com/gin-gonic/gin"
	v1handler "trungem.com/shopping-cart/internal/handler/v1"
)

type UserRoutes struct {
	handler *v1handler.UserHandler
}

func NewUserRoutes(handler *v1handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", ur.handler.GetAllUsers)
		users.POST("", ur.handler.CreateUser)
		users.GET("/:uuid", ur.handler.GetUserByUUID)
		users.PUT("/:uuid", ur.handler.UpdateUser)
		users.DELETE("/:uuid", ur.handler.DeleteUser)
	}
}
