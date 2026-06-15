package mail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"trungem.com/shopping-cart/internal/utils"
	"trungem.com/shopping-cart/pkg/logger"
)

type MailtrapConfig struct {
	MailSender     string
	NameSender     string
	MailTrapUrl    string
	MailTrapApiKey string
}

type MailtrapProvider struct {
	client *http.Client
	config *MailtrapConfig
	logger *zerolog.Logger
}

func NewMailtrapProvider(config *EmailConfig) (EmailProviderService, error) {
	mailtrapCfg, ok := config.ProviderConfig["mailtrap"].(map[string]any)
	if !ok {
		return nil, utils.NewError("Invalid or missing MailTrap configuration", utils.ErrCodeInternal)
	}

	return &MailtrapProvider{
		client: &http.Client{
			Timeout: config.Timeout,
		},
		config: &MailtrapConfig{
			MailSender:     mailtrapCfg["mail_sender"].(string),
			NameSender:     mailtrapCfg["name_sender"].(string),
			MailTrapUrl:    mailtrapCfg["mailtrap_url"].(string),
			MailTrapApiKey: mailtrapCfg["mailtrap_api_key"].(string),
		},
		logger: config.Logger,
	}, nil
}

func (p *MailtrapProvider) SendEmail(ctx context.Context, email *Email) error {
	traceId := logger.GetTraceID(ctx)
	start := time.Now()

	time.Sleep(5 * time.Second)

	email.From = Address{
		Email: p.config.MailSender,
		Name:  p.config.NameSender,
	}

	payload, err := json.Marshal(email)
	if err != nil {
		return utils.WrapError("Failed to marshal email", utils.ErrCodeInternal, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.config.MailTrapUrl, bytes.NewReader(payload))
	if err != nil {
		return utils.WrapError("Failed to create request", utils.ErrCodeInternal, err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.config.MailTrapApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := p.client.Do(req)
	if err != nil {
		p.logger.Error().Str("trace_id", traceId).Dur("duration", time.Since(start)).Str("operation", "send_mail").Err(err).Msg("Failed to send email")
		return utils.WrapError("Failed to send email", utils.ErrCodeInternal, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			p.logger.Error().Str("trace_id", traceId).Dur("duration", time.Since(start)).Str("operation", "send_mail").Err(err).Msg("Failed to close response body")
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		p.logger.Error().Str("trace_id", traceId).Dur("duration", time.Since(start)).Str("operation", "send_mail").Int("status_code", res.StatusCode).Str("response_body", string(body)).Msg("Failed to send email")
		return utils.NewError(fmt.Sprintf("Failed to send email, status code: %d", res.StatusCode), utils.ErrCodeInternal)
	}

	return nil
}
