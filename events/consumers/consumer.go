package consumers

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/yonraz/gochat_users/constants"
)

type Consumer struct {
	topic constants.Topic
	Consumer sarama.Consumer
}


func CreateConsumer(brokerList []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	
	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}
	
	return consumer, nil
}

func NewUserRegisteredConsumer(brokers []string) (*Consumer, error) {
	c, err := CreateConsumer(brokers)
	if err != nil {
		return nil, fmt.Errorf("error connecting consumer: %v", err)
	}
	return &Consumer{
		topic: constants.UserRegistered,
		Consumer: c,
	}, nil
}