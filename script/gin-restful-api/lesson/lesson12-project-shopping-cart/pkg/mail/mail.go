package mail

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"trungem.com/shopping-cart/internal/config"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/pkg/logger"
)

type Email struct {
	From     Address   `json:"from"`
	To       []Address `json:"to"`
	Subject  string    `json:"subject"`
	Text     string    `json:"text"`
	Category string    `json:"category"`
}

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type EmailConfig struct {
	ProviderConfig map[string]any
	ProviderType   ProviderType
	MaxRetries     int
	Timeout        time.Duration
	Logger         *zerolog.Logger
}

type EmailService struct {
	config   *EmailConfig
	provider EmailProviderService
	logger   *zerolog.Logger
}

func NewMailService(cfg *config.Config, logger *zerolog.Logger, providerFactory ProviderFactory) (EmailProviderService, error) {
	configEmail := &EmailConfig{
		ProviderConfig: cfg.MailProviderConfig,
		ProviderType:   ProviderType(cfg.MailProviderType),
		MaxRetries:     3,
		Timeout:        10 * time.Second,
		Logger:         logger,
	}

	provider, err := providerFactory.CreateProvider(configEmail)
	if err != nil {
		return nil, utils.WrapError("Failed to create provider", utils.ErrCodeInternal, err)
	}

	return &EmailService{
		config:   configEmail,
		provider: provider,
		logger:   logger,
	}, nil
}

func (ms *EmailService) SendEmail(ctx context.Context, email *Email) error {
	traceId := logger.GetTraceID(ctx)
	start := time.Now()

	var lastErr error
	for attempts := 1; attempts <= ms.config.MaxRetries; attempts++ {
		startAttempt := time.Now()
		err := ms.provider.SendEmail(ctx, email)
		if err == nil {
			ms.logger.Info().Str("trace_id", traceId).
				Dur("duration", time.Since(start)).
				Str("operation", "send_mail").
				Interface("to", email.To).
				Str("subject", email.Subject).
				Str("category", email.Category).
				Msg("Email sent successfully")
			return nil
		}

		lastErr = err
		ms.logger.Warn().Str("trace_id", traceId).
			Dur("duration", time.Since(startAttempt)).
			Str("operation", "send_mail_attempt").
			Int("attempt", attempts).
			Err(err).
			Msg("Failed to send email, retrying...")
		time.Sleep(time.Duration(attempts) * time.Second)
	}

	ms.logger.Error().Str("trace_id", traceId).
		Dur("duration", time.Since(start)).
		Str("operation", "send_mail").
		Err(lastErr).
		Msg("Failed to send email after max retries")

	return utils.WrapError("Failed to send email after all retries", utils.ErrCodeInternal, lastErr)
}
