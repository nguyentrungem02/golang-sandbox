package repository

import (
	"context"

	"github.com/google/uuid"
	"trungem.com/hoc-gin/internal/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
	FindByUuid(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
}
