package main

import (
	"context"
	"encoding/json"
	"errors"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/lpernett/godotenv"
	"github.com/rs/zerolog"
	"trungem.com/shopping-cart/internal/config"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/pkg/logger"
	"trungem.com/shopping-cart/pkg/mail"
	"trungem.com/shopping-cart/pkg/rabbitmq"
)

type Worker struct {
	rabbitMQ    rabbitmq.RabbitMQService
	mailService mail.EmailProviderService
	cfg         *config.Config
	logger      *zerolog.Logger
}

func NewWorker(cfg *config.Config) *Worker {
	log := utils.NewLoggerWithPath("worker.log", "info")

	// Connect RabbitMQ
	rabbitMQ, err := rabbitmq.NewRabbitMQService(utils.GetEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"), log)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init RabbitMQ service")
	}

	// Init mail service
	mailLogger := utils.NewLoggerWithPath("mail.log", "info")
	factory, err := mail.NewProviderFactory(mail.ProviderMailtrap)
	if err != nil {
		mailLogger.Error().Err(err).Msg("Error initializing mail provider")
		return nil
	}

	mailService, err := mail.NewMailService(cfg, mailLogger, factory)
	if err != nil {
		mailLogger.Error().Err(err).Msg("Error initializing mail service")
		return nil
	}

	return &Worker{
		rabbitMQ:    rabbitMQ,
		mailService: mailService,
		cfg:         cfg,
		logger:      log,
	}
}

func (w *Worker) Start(ctx context.Context) error {
	const emailQueueName = "auth_email_queue"

	handler := func(body []byte) error {
		w.logger.Debug().Msgf("Received email: %s", string(body))

		var email mail.Email
		if err := json.Unmarshal(body, &email); err != nil {
			w.logger.Error().Err(err).Msg("Failed to unmarshal email content")
			return err
		}

		if err := w.mailService.SendEmail(ctx, &email); err != nil {
			return utils.NewError("Failed to send password reset mail", utils.ErrCodeInternal)
		}

		w.logger.Info().Msgf("Email send successfully to: %v", email.To)

		return nil
	}

	if err := w.rabbitMQ.Consume(ctx, emailQueueName, handler); err != nil {
		w.logger.Error().Err(err).Msg("Failed to start consumer")
		return err
	}

	w.logger.Info().Msgf("Worker started, consuming from queue: %s,", emailQueueName)

	<-ctx.Done()
	w.logger.Info().Msg("Worker stopped, consuming due to connect cancellation")

	return ctx.Err()
}

func (w *Worker) Shutdown(ctx context.Context) error {
	w.logger.Info().Msg("Worker shutting down ...")

	if err := w.rabbitMQ.Close(); err != nil {
		w.logger.Error().Err(err).Msg("Failed to close RabbitMQ connection")
		return err
	}

	w.logger.Info().Msg("RabbitMQ connection closed successfully")

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			w.logger.Warn().Msg("Worker shutdown timed out")
			return ctx.Err()
		}
	default:
	}

	return nil
}

func main() {
	rootDir := utils.MustGetWorkingDir()

	logFile := filepath.Join(rootDir, "internal/logs/app.log")

	logger.InitLogger(logger.LoggerConfig{
		Level:      "info",
		Filename:   logFile,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
		IsDev:      utils.GetEnv("APP_ENV", "development"),
	})

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		logger.Log.Warn().Msg("⚠️ No .env file found")
	} else {
		logger.Log.Info().Msg("✅ Loaded successfully .env in worker process")
	}

	// Initialize configuration
	cfg := config.NewConfig()

	worker := NewWorker(cfg)
	if worker == nil {
		logger.Log.Fatal().Msg("Failed to initialize worker")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := worker.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
			logger.Log.Fatal().Err(err).Msg("Worker stopped with error")
			return
		}
	}()

	<-ctx.Done()
	logger.Log.Warn().Msg("Received shutdown signal, stopping worker...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := worker.Shutdown(shutdownCtx); err != nil {
		logger.Log.Error().Err(err).Msg("Worker shutdown error")
	}

	wg.Wait()
	logger.Log.Info().Msg("Main process terminated")
}
