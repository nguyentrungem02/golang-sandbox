package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1dto "trungem.com/shopping-cart/internal/dto/v1"
	v1service "trungem.com/shopping-cart/internal/service/v1"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/internal/validation"
)

type AuthHandler struct {
	service v1service.AuthService
}

func NewAuthHandler(service v1service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var input v1dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	userLogin, accessToken, refreshToken, expiresIn, err := ah.service.Login(ctx, input.Email, input.Password)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(userLogin)

	response := utils.NewTokenResponse(userDTO, accessToken, refreshToken, expiresIn)

	utils.ResponseSuccess(ctx, http.StatusOK, "Login successfully", response)
}

func (ah *AuthHandler) Logout(ctx *gin.Context) {
	var input v1dto.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	if err := ah.service.Logout(ctx, input.RefreshToken); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Logout successfully")
}

func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	var input v1dto.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	user, accessToken, refreshToken, expiresIn, err := ah.service.RefreshToken(ctx, input.RefreshToken)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(user)

	response := utils.NewTokenResponse(userDTO, accessToken, refreshToken, expiresIn)

	utils.ResponseSuccess(ctx, http.StatusOK, "Refresh token generate successfully", response)
}

func (ah *AuthHandler) RequestForgotPassword(ctx *gin.Context) {
	var input v1dto.RequestPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	if err := ah.service.RequestForgotPassword(ctx, input.Email); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Reset link sent to email")
}

func (ah *AuthHandler) ResetPassword(ctx *gin.Context) {
	var input v1dto.ResetPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidation(ctx, validation.HandleValidationErrors(err))
		return
	}

	if err := ah.service.ResetPassword(ctx, input.Token, input.NewPassword); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Password reset successfully")
}
