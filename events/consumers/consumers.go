package consumers

import (
	"github.com/streadway/amqp"
	"github.com/yonraz/gochat_users/constants"
)

func NewUserRegisteredConsumer(channel *amqp.Channel) *Consumer {
	return NewConsumer(channel, constants.UserRegistrationQueue, constants.UserRegisteredKey, constants.UserEventsExchange, UserRegisteredHandler)
}

func NewUserLoggedinConsumer(channel *amqp.Channel) *Consumer {
	return NewConsumer(channel, constants.UserLoginQueue, constants.UserLoggedInKey, constants.UserEventsExchange, UserLoggedinHandler)
}

func NewUserSignedoutConsumer(channel *amqp.Channel) *Consumer {
	return NewConsumer(channel, constants.UserSignoutQueue, constants.UserSignedoutKey, constants.UserEventsExchange, UserSignedoutHandler)
}