package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"github.com/yonraz/gochat_users/controllers"
	"github.com/yonraz/gochat_users/events/consumers"
	"github.com/yonraz/gochat_users/initializers"
	// "github.com/yonraz/gochat_users/services"
)

func init () {
	fmt.Println("Application starting...")
	time.Sleep(1 * time.Minute)
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ConnectToRabbitmq()
	initializers.ConnectToRedis()
}

func main() {
	router := gin.Default()

	defer func() {
		if err := initializers.RmqChannel.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ channel: %v", err)
		}
	}()
	defer func() {
		if err := initializers.RmqConn.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
	}()
	usersController := controllers.NewUsersController()

	// dev - insert 35 lines of mock user data
	// services.CreateMockUsers()
	//

	go startConsumers()

	router.GET("/api/users", usersController.GetUsers)

	router.Run()
}

func startConsumers() {
	targetConsumers := []func(*amqp.Channel) *consumers.RmqConsumer{
		consumers.NewUserRegisteredConsumer,
		consumers.NewUserLoggedinConsumer,
		consumers.NewUserSignedoutConsumer,
	}

	for _, consumerFunc := range targetConsumers {
		consumer := consumerFunc(initializers.RmqChannel)
		go func(c *consumers.RmqConsumer) {
			if err := c.Consume(); err != nil {
				log.Fatalf("Error starting consumer: %v", err)
			}
		}(consumer)
	}
}