package configs

import (
	"fmt"

	"github.com/streadway/amqp"
)

func NewRabbitMQConnection() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://" + rabbitmqUser + ":" + rabbitmqPass + "@" + rabbitmqHost + ":5672")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return conn, ch, nil
}
