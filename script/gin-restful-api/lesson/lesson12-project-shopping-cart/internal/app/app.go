package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"trungem.com/shopping-cart/internal/config"
	"trungem.com/shopping-cart/internal/db"
	"trungem.com/shopping-cart/internal/db/sqlc"
	"trungem.com/shopping-cart/internal/middleware"
	"trungem.com/shopping-cart/internal/routes"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/internal/validation"
	"trungem.com/shopping-cart/pkg/auth"
	"trungem.com/shopping-cart/pkg/cache"
	"trungem.com/shopping-cart/pkg/logger"
	"trungem.com/shopping-cart/pkg/mail"
	"trungem.com/shopping-cart/pkg/rabbitmq"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

type ModuleContext struct {
	DB    sqlc.Querier
	Redis *redis.Client
}

func NewApplication(cfg *config.Config) (*Application, error) {
	r := gin.Default()

	if err := validation.InitValidator(); err != nil {
		logger.Log.Fatal().Err(err).Msg("Error initializing validator")
	}

	go middleware.CleanupClients()

	if err := db.InitDB(); err != nil {
		logger.Log.Fatal().Err(err).Msg("Database init failed")
		return nil, err
	}

	redisClient := config.NewRedisClient()
	cacheRedisService := cache.NewRedisCacheService(redisClient)
	tokenService := auth.NewJWTService(cacheRedisService)
	mailLogger := utils.NewLoggerWithPath("mail.log", "info")
	factory, err := mail.NewProviderFactory(mail.ProviderMailtrap)
	if err != nil {
		mailLogger.Error().Err(err).Msg("Error initializing mail provider")
		return nil, err
	}

	mailService, err := mail.NewMailService(cfg, mailLogger, factory)
	if err != nil {
		mailLogger.Error().Err(err).Msg("Error initializing mail service")
		return nil, err
	}

	rabbitmqLogger := utils.NewLoggerWithPath("worker.log", "info")
	rabbitmqService, _ := rabbitmq.NewRabbitMQService(utils.GetEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"), rabbitmqLogger)

	ctx := &ModuleContext{
		DB:    db.DB,
		Redis: redisClient,
	}

	module := []Module{
		NewUserModule(ctx),
		NewAuthModule(ctx, tokenService, cacheRedisService, mailService, rabbitmqService),
	}

	routes.RegisterRoutes(r, tokenService, cacheRedisService, getModuleRoutes(module)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: module,
	}, nil
}

func (a *Application) Run() error {
	srv := &http.Server{
		Addr:    a.config.ServerAddress,
		Handler: a.router,
	}

	quit := make(chan os.Signal, 1)
	// syscall.SIGINT -> Ctrl + C
	// syscall.SIGTERM -> Kill service
	// syscall.SIGHUP -> Reload service
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		logger.Log.Info().Msgf("✅ Server is running at %s", srv.Addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Fatal().Err(err).Msg("⛔ Failed to start server")
		}
	}()

	<-quit
	logger.Log.Warn().Msg("️ Shutdown signal received ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("⛔ Server forced to shutdown")
	}

	logger.Log.Info().Msg("🍺 Server exited gracefully")

	return nil
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}
	return routeList
}
