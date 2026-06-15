package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"trungem.com/hoc-gin/internal/db/sqlc"
	"trungem.com/hoc-gin/internal/repository"
)

type UserResponse struct {
	UserId    int32     `json:"user_id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
}

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")

	userUuid, err := uuid.Parse(uuidParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	user, err := uh.repo.FindByUuid(ctx, userUuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := UserResponse{
		UserId:    user.UserID,
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input sqlc.CreateUserParams
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.repo.Create(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := UserResponse{
		UserId:    user.UserID,
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": response,
	})
}
