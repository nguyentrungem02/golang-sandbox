package v1service

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
	"trungem.com/shopping-cart/internal/db/sqlc"
	"trungem.com/shopping-cart/internal/repository"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/pkg/auth"
	"trungem.com/shopping-cart/pkg/cache"
	"trungem.com/shopping-cart/pkg/logger"
	"trungem.com/shopping-cart/pkg/mail"
	"trungem.com/shopping-cart/pkg/rabbitmq"
)

type authService struct {
	userRepo     repository.UserRepository
	tokenService auth.TokenService
	cacheService cache.RedisCacheService
	mailService  mail.EmailProviderService
	rabbitmq     rabbitmq.RabbitMQService
}

type LoginAttempt struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu               sync.Mutex
	clients          = make(map[string]*LoginAttempt)
	LoginAttemptTTL  = 5 * time.Minute
	MaxLoginAttempts = 5
)

func NewAuthService(userRepo repository.UserRepository, tokenService auth.TokenService, cacheService cache.RedisCacheService, mailService mail.EmailProviderService, rabbitmqService rabbitmq.RabbitMQService) AuthService {
	return &authService{
		userRepo:     userRepo,
		tokenService: tokenService,
		cacheService: cacheService,
		mailService:  mailService,
		rabbitmq:     rabbitmqService,
	}
}

func (as *authService) getClientIp(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func (as *authService) getLoginAttempt(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Limit(float32(MaxLoginAttempts)/float32(LoginAttemptTTL.Seconds())), MaxLoginAttempts)
		newClient := &LoginAttempt{limiter, time.Now()}
		clients[ip] = newClient

		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func (as *authService) checkLoginAttempt(ip string) error {
	limiter := as.getLoginAttempt(ip)

	if !limiter.Allow() {
		return utils.NewError("Too many login attempts. Please retry again later", utils.ErrCodeTooManyRequests)
	}

	return nil
}

func (as *authService) CleanupClients(ip string) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, ip)
}

func (as *authService) Login(ctx *gin.Context, email, password string) (sqlc.User, string, string, int, error) {
	context := ctx.Request.Context()
	ip := as.getClientIp(ctx)

	if err := as.checkLoginAttempt(ip); err != nil {
		return sqlc.User{}, "", "", 0, err
	}

	email = utils.NormalizeString(email)
	user, err := as.userRepo.GetByEmail(context, email)
	if err != nil {
		as.getLoginAttempt(ip)
		return sqlc.User{}, "", "", 0, utils.NewError("Invalid email or password", utils.ErrCodeUnauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password)); err != nil {
		as.getLoginAttempt(ip)
		return sqlc.User{}, "", "", 0, utils.NewError("Invalid email or password", utils.ErrCodeUnauthorized)
	}

	accessToken, err := as.tokenService.GenerateAccessToken(user)
	if err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Unable to create access token", utils.ErrCodeInternal, err)
	}

	refreshToken, err := as.tokenService.GenerateRefreshToken(user)
	if err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Unable to create refresh token", utils.ErrCodeInternal, err)
	}

	if err := as.tokenService.StoreRefreshToken(refreshToken); err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Cannot save refresh token", utils.ErrCodeInternal, err)
	}

	as.CleanupClients(ip)

	return user, accessToken, refreshToken.Token, int(auth.AccessTokenTTL.Seconds()), nil
}

func (as *authService) Logout(ctx *gin.Context, refreshToken string) error {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.NewError("Missing Authorization header", utils.ErrCodeUnauthorized)
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	_, claims, err := as.tokenService.ParseToken(accessToken)
	if err != nil {
		return utils.NewError("Invalid access token", utils.ErrCodeUnauthorized)
	}

	if jti, ok := claims["jti"].(string); ok {
		expUnix, _ := claims["exp"].(float64)
		exp := time.Unix(int64(expUnix), 0)
		key := "blacklist:" + jti
		ttl := time.Until(exp)

		if err := as.cacheService.Set(key, "revoked", ttl); err != nil {
			return utils.WrapError("Cannot save access token to backlist", utils.ErrCodeInternal, err)
		}
	}

	_, err = as.tokenService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return utils.WrapError("Refresh token is invalid or revoked", utils.ErrCodeUnauthorized, err)
	}

	if err = as.tokenService.RevokeRefreshToken(refreshToken); err != nil {
		return utils.WrapError("Unable to revoke refresh token", utils.ErrCodeInternal, err)
	}

	return nil
}

func (as *authService) RefreshToken(ctx *gin.Context, token string) (sqlc.User, string, string, int, error) {
	context := ctx.Request.Context()

	refreshToken, err := as.tokenService.ValidateRefreshToken(token)
	if err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Refresh token is invalid or revoked", utils.ErrCodeUnauthorized, err)
	}

	userUuid, _ := uuid.Parse(refreshToken.UserUUID)
	user, err := as.userRepo.FindByUUID(context, userUuid)
	if err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("User not found", utils.ErrCodeUnauthorized, err)
	}

	accessTokenNew, err := as.tokenService.GenerateAccessToken(user)
	if err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Unable to create access token", utils.ErrCodeInternal, err)
	}

	refreshTokenNew, err := as.tokenService.GenerateRefreshToken(user)
	if err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Unable to create refresh token", utils.ErrCodeInternal, err)
	}

	if err := as.tokenService.RevokeRefreshToken(token); err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Unable to revoke refresh token", utils.ErrCodeInternal, err)
	}

	if err := as.tokenService.StoreRefreshToken(refreshTokenNew); err != nil {
		return sqlc.User{}, "", "", 0, utils.WrapError("Cannot save refresh token", utils.ErrCodeInternal, err)
	}

	return user, accessTokenNew, refreshTokenNew.Token, int(auth.AccessTokenTTL.Seconds()), nil
}

func (as *authService) RequestForgotPassword(ctx *gin.Context, email string) error {
	context := ctx.Request.Context()

	rateLimitKey := fmt.Sprintf("reset:ratelimit:%s", email)
	if exists, err := as.cacheService.Exists(rateLimitKey); err == nil && exists {
		return utils.NewError("Please wait before requesting another password reset", utils.ErrCodeTooManyRequests)
	}

	user, err := as.userRepo.GetByEmail(context, email)
	if err != nil {
		return utils.NewError("Email not found", utils.ErrCodeNotFound)
	}

	token, err := utils.GenerateRandomString(16)
	if err != nil {
		return utils.NewError("Failed to generate token", utils.ErrCodeInternal)
	}

	if err = as.cacheService.Set("reset:"+token, user.UserUuid, time.Hour); err != nil {
		return utils.NewError("Failed to store reset token", utils.ErrCodeInternal)
	}

	if err = as.cacheService.Set(rateLimitKey, "1", 5*time.Minute); err != nil {
		return utils.NewError("Failed to store rate limit reset password", utils.ErrCodeInternal)
	}

	resetLink := fmt.Sprintf("https://abc.com/view-to-reset-password?token=%s", token)

	logger.Log.Info().Msg(resetLink)

	mailContent := &mail.Email{
		To: []mail.Address{
			{Email: email},
		},
		Subject: "Password Reset Request",
		Text:    fmt.Sprintf("Hi %s,\n\nYou requested a password reset. Click the link below to reset your password:\n\n%s\n\nIf you didn't request this, please ignore this email.", user.UserEmail, resetLink),
	}

	if err = as.rabbitmq.Publish(ctx, "auth_email_queue", mailContent); err != nil {
		return utils.NewError("Failed to send password reset mail", utils.ErrCodeInternal)
	}

	//if err = as.mailService.SendEmail(context, mailContent); err != nil {
	//	return utils.NewError("Failed to send password reset mail", utils.ErrCodeInternal)
	//}

	return nil
}

func (as *authService) ResetPassword(ctx *gin.Context, token, password string) error {
	context := ctx.Request.Context()

	var userUUIDString string
	err := as.cacheService.Get("reset:"+token, &userUUIDString)
	if errors.Is(err, redis.Nil) || userUUIDString == "" {
		return utils.NewError("Invalid or expired token", utils.ErrCodeNotFound)
	}

	if err != nil {
		return utils.NewError("Failed to get reset token", utils.ErrCodeInternal)
	}

	userUuid, err := uuid.Parse(userUUIDString)
	if err != nil {
		return utils.WrapError("Uuid is invalid", utils.ErrCodeInternal, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return utils.WrapError("Failed to hash password", utils.ErrCodeInternal, err)
	}

	var input = sqlc.UpdatePasswordParams{
		UserPassword: string(hashedPassword),
		UserUuid:     userUuid,
	}
	_, err = as.userRepo.UpdatePassword(context, input)
	if err != nil {
		return utils.NewError("Failed to update password", utils.ErrCodeInternal)
	}

	if err = as.cacheService.Clear("reset:" + token); err != nil {
		return utils.NewError("Failed to clear cache", utils.ErrCodeInternal)
	}

	return nil
}
