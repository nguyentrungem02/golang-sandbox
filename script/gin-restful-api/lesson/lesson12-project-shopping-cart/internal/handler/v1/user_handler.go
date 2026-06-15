package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	users, total, err := uh.service.GetAllUsers(ctx, params.Search, params.Order, params.Sort, params.Page, params.Limit, false)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	usersDTO := v1dto.MapUsersToDTO(users)

	paginationResp := utils.NewPaginationResponse(usersDTO, params.Page, params.Limit, total)

	utils.ResponseSuccess(ctx, http.StatusOK, "User list successfully", paginationResp)
}

func (uh *UserHandler) GetAllUsersSoftDeleted(ctx *gin.Context) {
	var params v1dto.GetAllUsersSoftDeletedParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	usersSoftDeleted, total, err := uh.service.GetAllUsers(ctx, params.Search, params.Order, params.Sort, params.Page, params.Limit, true)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	usersSoftDeletedDTO := v1dto.MapUsersToDTO(usersSoftDeleted)

	paginationResp := utils.NewPaginationResponse(usersSoftDeletedDTO, params.Page, params.Limit, total)

	utils.ResponseSuccess(ctx, http.StatusOK, "User list soft deleted successfully", paginationResp)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input v1dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	user := input.MapCreateInputToModel()

	newUser, err := uh.service.CreateUser(ctx, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(newUser)

	utils.ResponseSuccess(ctx, http.StatusCreated, "User created successfully", userDTO)
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userByUuid, err := uh.service.GetUserByUUID(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(userByUuid)

	utils.ResponseSuccess(ctx, http.StatusOK, "Get user by uuid successfully", userDTO)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	var input v1dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	user := input.MapUpdateInputToModel(userUuid)

	updatedUser, err := uh.service.UpdateUser(ctx, user)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(updatedUser)

	utils.ResponseSuccess(ctx, http.StatusOK, "User updated successfully", userDTO)
}

func (uh *UserHandler) SoftDeleteUser(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	softDeletedUser, err := uh.service.SoftDeleteUser(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(softDeletedUser)

	utils.ResponseSuccess(ctx, http.StatusOK, "User soft deleted successfully", userDTO)
}

func (uh *UserHandler) RestoreUser(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	restoredUser, err := uh.service.RestoreUser(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(restoredUser)

	utils.ResponseSuccess(ctx, http.StatusOK, "User restored successfully", userDTO)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var param v1dto.UuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	err = uh.service.DeleteUser(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatusCode(ctx, http.StatusNoContent)
}
