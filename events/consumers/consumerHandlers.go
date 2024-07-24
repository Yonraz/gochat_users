package consumers

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/yonraz/gochat_users/controllers"
)

func UserRegisteredHandler(msg amqp.Delivery) error {
	fmt.Printf("User registered with data %v\n", msg.Body)
	var event map[string]string
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	username, ok := event["username"]
	if !ok {
		return fmt.Errorf("username field is missing in the message")
	}
	if err := controllers.Signup(username); err != nil {
		return err
	}
	fmt.Printf("User Registered: %s\n", username)
	return nil
}

func UserLoggedinHandler(msg amqp.Delivery) error {
	fmt.Printf("User logged in with data %v\n", msg.Body)

	return nil
}

func UserSignedoutHandler(msg amqp.Delivery) error {
	fmt.Printf("User signed out with data %v\n", msg.Body)

	return nil
}