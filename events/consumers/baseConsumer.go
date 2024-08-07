package consumers

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/yonraz/gochat_users/constants"
)

type Consumer struct {
	channel *amqp.Channel
	queueName string
	routingKey string
	exchange string
	handlerFunc func(amqp.Delivery) error
}

func NewConsumer(channel *amqp.Channel, queueName constants.Queues, routingKey constants.RoutingKey, exchange constants.Exchange, handlerFunc func(amqp.Delivery) error) *Consumer {
	return &Consumer{
		channel:     channel,
		queueName:   string(queueName),
		routingKey:  string(routingKey),
		exchange:    string(exchange),
		handlerFunc: handlerFunc,
	}
}

func (c *Consumer) Consume() error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming %w", err)
	}

	go func() {
		for msg := range msgs {
			if err := c.handlerFunc(msg); err != nil {
				log.Printf("Failed to process message: %v", err)
				// Handle nack or requeue logic if necessary
				msg.Nack(false, false)
			} else {
				msg.Ack(false)
			}
		}
	}()

	log.Printf("Started consuming on queue: %s", c.queueName)
	return nil
}

