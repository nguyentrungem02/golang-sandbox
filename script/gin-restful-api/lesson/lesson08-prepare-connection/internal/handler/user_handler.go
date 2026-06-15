package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"trungem.com/hoc-gin/internal/models"
	"trungem.com/hoc-gin/internal/repository"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	uh.repo.FindById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User By UUID",
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uh.repo.Create(user)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create User",
	})
}
