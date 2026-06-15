package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1dto "trungem.com/shopping-cart/internal/dto/v1"
	v1service "trungem.com/shopping-cart/internal/service/v1"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/internal/validation"
)

type UserHandler struct {
	service v1service.UserService
}

func NewUserHandler(service v1service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params v1dto.GetAllUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input v1dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "")
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	var input v1dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	utils.ResponseStatusCode(ctx, http.StatusNoContent)
}
