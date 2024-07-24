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
)

func init () {
	time.Sleep(15 * time.Second)
	fmt.Println("Application starting...")
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ConnectToRabbitmq()
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

	go startConsumers()

	router.GET("/api/users", controllers.GetUsers)

	router.Run()
}

func startConsumers() {
	targetConsumers := []func(*amqp.Channel) *consumers.Consumer{
		consumers.NewUserRegisteredConsumer,
		consumers.NewUserLoggedinConsumer,
		consumers.NewUserSignedoutConsumer,
	}

	for _, consumerFunc := range targetConsumers {
		consumer := consumerFunc(initializers.RmqChannel)
		go func(c *consumers.Consumer) {
			if err := c.Consume(); err != nil {
				log.Fatalf("Error starting consumer: %v", err)
			}
		}(consumer)
	}
}