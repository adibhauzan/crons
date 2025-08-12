package broker

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQProducerInterface interface {
	Publish(queue string, message []byte) error
}

type RabbitMQProducer struct {
	Channel *amqp.Channel
}

func NewRabbitMQProducer(conn *amqp.Connection) (*RabbitMQProducer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}
	return &RabbitMQProducer{
		Channel: ch,
	}, nil
}

func (p *RabbitMQProducer) Publish(queue string, message []byte) error {
	_, err := p.Channel.QueueDeclare(
		queue, // queue name
		true,  // durable (survive server restart)
		false, // auto-delete (delete when unused)
		false, // exclusive (only the current connection can access)
		false, // no-wait (don't wait for acknowledgment)
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", queue, err)
	}

	err = p.Channel.Publish(
		"",    // exchange (empty string for default)
		queue, // routing key (queue name)
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json", // could be parameterized
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}
