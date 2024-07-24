package consumers

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/yonraz/gochat_users/constants"
	"github.com/yonraz/gochat_users/models"
	"github.com/yonraz/gochat_users/services"
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
	user := models.User{Username: username, Status: constants.Offline}
	if err := services.Create(user); err != nil {
		return err
	}
	fmt.Printf("User Registered: %s\n", username)
	return nil
}

func UserLoggedinHandler(msg amqp.Delivery) error {
	var event map[string]string
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	username, ok := event["username"]
	if !ok {
		return fmt.Errorf("username field is missing in the message")
	}
	user := models.User{Username: username, Status: constants.Offline}
	if err := services.Login(user); err != nil {
		return err
	}
	fmt.Printf("User logged in: %s\n", username)
	return nil
}

func UserSignedoutHandler(msg amqp.Delivery) error {
	var event map[string]string
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	username, ok := event["username"]
	if !ok {
		return fmt.Errorf("username field is missing in the message")
	}
	user := models.User{Username: username, Status: constants.Offline}
	if err := services.Signout(user); err != nil {
		return err
	}
	fmt.Printf("User signed out: %s\n", username)
	return nil
}