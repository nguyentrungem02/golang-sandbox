package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trungem.com/user-manager/internal/dto"
	"trungem.com/user-manager/internal/service"
	"trungem.com/user-manager/internal/utils"
	"trungem.com/user-manager/internal/validation"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params dto.GetAllUsersParams
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

	users, err := uh.service.GetAllUsers(params.Search, params.Page, params.Limit)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDto := dto.MapUsersToDTO(users)

	utils.ResponseSuccess(ctx, http.StatusOK, &userDto)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	user := input.MapCreateInputToModel()

	createdUser, err := uh.service.CreateUser(user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDto := dto.MapUserToDTO(createdUser)

	utils.ResponseSuccess(ctx, http.StatusCreated, &userDto)
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var param dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	user, err := uh.service.GetUserByUUID(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDto := dto.MapUserToDTO(user)

	utils.ResponseSuccess(ctx, http.StatusOK, &userDto)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var param dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	var input dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	user := input.MapUpdateInputToModel()

	updatedUser, err := uh.service.UpdateUser(param.Uuid, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDto := dto.MapUserToDTO(updatedUser)

	utils.ResponseSuccess(ctx, http.StatusOK, &userDto)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var param dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	err := uh.service.DeleteUser(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatusCode(ctx, http.StatusNoContent)
}
