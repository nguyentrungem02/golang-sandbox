package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type rabbitMQService struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	logger  *zerolog.Logger
}

func NewRabbitMQService(amqpURL string, logger *zerolog.Logger) (RabbitMQService, error) {
	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to RabbitMQ")
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		if err = conn.Close(); err != nil {
			logger.Error().Err(err).Msg("Failed to close RabbitMQ connection after channel creation failure")
			return nil, err
		}

		logger.Error().Err(err).Msg("Failed to open channel")
		return nil, err
	}

	return &rabbitMQService{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *rabbitMQService) Publish(ctx context.Context, queue string, message any) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to declare a queue")
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to marshal message")
		return err
	}

	err = r.channel.PublishWithContext(ctx, "", queue, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to publish a message")
		return err
	}

	return nil
}

func (r *rabbitMQService) Consume(ctx context.Context, queue string, handler func([]byte) error) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to declare a queue")
		return err
	}

	messages, err := r.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to register a consumer")
		return err
	}

	go func() {
		for {
			select {
			case msg, ok := <-messages:
				if !ok {
					return
				}

				if err := handler(msg.Body); err != nil {
					if err = msg.Nack(false, false); err != nil {
						return
					}
				} else {

					if err = msg.Ack(false); err != nil {
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (r *rabbitMQService) Close() error {
	if r.channel != nil {
		if err := r.channel.Close(); err != nil {
			r.logger.Error().Err(err).Msg("Failed to close RabbitMQ channel")
			return err
		}
	}

	if r.conn != nil {
		if err := r.conn.Close(); err != nil {
			r.logger.Error().Err(err).Msg("Failed to close RabbitMQ connection")
			return err
		}
	}

	return nil
}
