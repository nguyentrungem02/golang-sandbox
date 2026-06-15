package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trungem.com/hoc-golang/utils"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all users (V1)",
	})
}

func (u *UserHandler) GetUsersByIdV1(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := utils.ValidationPositiveInt("ID", idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by id (V1)",
		"id":      id,
	})
}

func (u *UserHandler) GetUsersByUuidV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")

	uid, err := utils.ValidationUuid("UUID", uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by uuid (V1)",
		"uuid":    uid,
	})
}

func (u *UserHandler) PostUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user (V1)",
	})
}

func (u *UserHandler) PutUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user (V1)",
	})
}

func (u *UserHandler) DeleteUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete user (V1)",
	})
}
