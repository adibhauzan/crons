package utils

import (
	"encoding/json"
	"fmt"

	"github.com/adibhauzan/crons/internal/broker"
)

func PublishTasks(broker broker.RabbitMQProducerInterface, queueName string, data interface{}) error {
	messageBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := broker.Publish(queueName, messageBytes); err != nil {
		return fmt.Errorf("failed to publish to RabbitMQ: %w", err)
	}

	return nil
}
