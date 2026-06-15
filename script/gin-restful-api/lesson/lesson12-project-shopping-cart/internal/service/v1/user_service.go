package v1service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"trungem.com/shopping-cart/internal/db/sqlc"
	"trungem.com/shopping-cart/internal/repository"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/pkg/cache"
	"trungem.com/shopping-cart/pkg/logger"
)

type userService struct {
	repo  repository.UserRepository
	cache cache.RedisCacheService
}

func NewUserService(repo repository.UserRepository, redisClient *redis.Client) UserService {
	return &userService{
		repo:  repo,
		cache: cache.NewRedisCacheService(redisClient),
	}
}

func (us *userService) GetAllUsers(ctx *gin.Context, search, orderBy, sort string, page, limit int32, deleted bool) ([]sqlc.User, int32, error) {
	context := ctx.Request.Context()

	/** Get Cache Redis **/
	cacheKey := us.generateCacheKey(search, orderBy, sort, page, limit, deleted)
	var cacheData struct {
		Users []sqlc.User `json:"users"`
		Total int32       `json:"total"`
	}

	if err := us.cache.Get(cacheKey, &cacheData); err == nil && cacheData.Users != nil {
		return cacheData.Users, cacheData.Total, nil
	}

	if sort == "" {
		sort = "desc"
	}

	if orderBy == "" {
		orderBy = "user_created_at"
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limitInt := utils.GetIntEnv("LIMIT_ITEM_ON_PER_PAGE", 10)
		limit = int32(limitInt)
	}

	offset := (page - 1) * limit

	users, err := us.repo.GetAllV2(context, search, orderBy, sort, limit, offset, deleted)
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapError("failed to fetch users", utils.ErrCodeInternal, err)
	}

	total, err := us.repo.CountAll(ctx, search, deleted)
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapError("failed to count users", utils.ErrCodeInternal, err)
	}

	//Create cache data
	cacheData = struct {
		Users []sqlc.User `json:"users"`
		Total int32       `json:"total"`
	}{
		Users: users,
		Total: int32(total),
	}

	if err := us.cache.Set(cacheKey, cacheData, 5*time.Minute); err != nil {
		logger.Log.Err(err).Msg("Failed cache data redis: %v")
	}

	return users, int32(total), nil
}

func (us *userService) CreateUser(ctx *gin.Context, input sqlc.CreateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()

	input.UserEmail = utils.NormalizeString(input.UserEmail)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, utils.WrapError("failed to hash password", utils.ErrCodeInternal, err)
	}

	input.UserPassword = string(hashPassword)

	newUser, err := us.repo.Create(context, input)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return sqlc.User{}, utils.WrapError("email already exist", utils.ErrCodeConflict, err)
		}

		return sqlc.User{}, utils.WrapError("failed to create a new user", utils.ErrCodeInternal, err)
	}

	// Clear cache
	if err := us.cache.Clear("users:*"); err != nil {
		logger.Log.Warn().Err(err).Msg("Failed clear cache data redis")
	}

	return newUser, nil
}

func (us *userService) GetUserByUUID(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()

	user, err := us.repo.FindByUUID(context, uuid)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return sqlc.User{}, utils.WrapError("user does not exist", utils.ErrCodeNotFound, err)
		}

		return sqlc.User{}, utils.WrapError("failed to fetch user by uuid", utils.ErrCodeInternal, err)
	}

	return user, nil
}

func (us *userService) UpdateUser(ctx *gin.Context, input sqlc.UpdateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()

	if input.UserPassword != nil && *input.UserPassword != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(*input.UserPassword), bcrypt.DefaultCost)
		if err != nil {
			return sqlc.User{}, utils.WrapError("failed to hash password", utils.ErrCodeInternal, err)
		}

		hashed := string(hashPassword)
		input.UserPassword = &hashed
	}

	updatedUser, err := us.repo.Update(context, input)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return sqlc.User{}, utils.WrapError("user does not exist", utils.ErrCodeNotFound, err)
		}
		return sqlc.User{}, utils.WrapError("failed to update user", utils.ErrCodeInternal, err)
	}

	// Clear cache
	if err := us.cache.Clear("users:*"); err != nil {
		logger.Log.Err(err).Msg("Failed clear cache data redis")
	}

	return updatedUser, nil
}

func (us *userService) SoftDeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()

	softDeletedUser, err := us.repo.SoftDelete(context, uuid)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return sqlc.User{}, utils.WrapError("user does not exist", utils.ErrCodeNotFound, err)
		}

		return sqlc.User{}, utils.WrapError("failed to delete user", utils.ErrCodeInternal, err)
	}

	// Clear cache
	if err := us.cache.Clear("users:*"); err != nil {
		logger.Log.Err(err).Msg("Failed clear cache data redis")
	}

	return softDeletedUser, nil
}

func (us *userService) RestoreUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()

	restoredUser, err := us.repo.Restore(context, uuid)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return sqlc.User{}, utils.WrapError("user does not exist or marked as delete for restore", utils.ErrCodeNotFound, err)
		}

		return sqlc.User{}, utils.WrapError("failed to restore user", utils.ErrCodeInternal, err)
	}

	// Clear cache
	if err := us.cache.Clear("users:*"); err != nil {
		logger.Log.Err(err).Msg("Failed clear cache data redis")
	}

	return restoredUser, nil
}

func (us *userService) DeleteUser(ctx *gin.Context, uuid uuid.UUID) error {
	context := ctx.Request.Context()

	err := us.repo.Delete(context, uuid)

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return utils.WrapError("user does not exist or marked as delete for permanent removal", utils.ErrCodeNotFound, err)
		}

		return utils.WrapError("failed to delete user", utils.ErrCodeInternal, err)
	}

	// Clear cache
	if err := us.cache.Clear("users:*"); err != nil {
		logger.Log.Err(err).Msg("Failed clear cache data redis")
	}

	return nil
}

func (us *userService) generateCacheKey(search, orderBy, sort string, page, limit int32, deleted bool) string {
	search = strings.TrimSpace(search)
	if search == "" {
		search = "none"
	}

	orderBy = strings.TrimSpace(orderBy)
	if orderBy == "" {
		orderBy = "user_created_at"
	}

	sort = strings.ToLower(strings.TrimSpace(sort))
	if sort == "" {
		sort = "desc"
	}

	return fmt.Sprintf("users:%s:%s:%s:%d:%d:%t", search, orderBy, sort, page, limit, deleted)
}
