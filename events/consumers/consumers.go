package consumers

import (
	"github.com/streadway/amqp"
	"github.com/yonraz/gochat_users/constants"
)

// func NewUserRegisteredConsumer(channel *amqp.Channel) *RmqConsumer {
// 	return NewConsumer(channel, constants.UserRegistrationQueue, constants.UserRegistered, constants.UserEventsExchange, UserRegisteredHandler)
// }

// func NewUserLoggedinConsumer(channel *amqp.Channel) *RmqConsumer {
// 	return NewConsumer(channel, constants.UserLoginQueue, constants.UserLoggedIn, constants.UserEventsExchange, UserLoggedinHandler)
// }

// func NewUserSignedoutConsumer(channel *amqp.Channel) *RmqConsumer {
// 	return NewConsumer(channel, constants.UserSignoutQueue, constants.UserSignedout, constants.UserEventsExchange, UserSignedoutHandler)
// }