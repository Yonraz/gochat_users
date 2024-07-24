package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yonraz/gochat_users/initializers"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	
	router := gin.Default()

	router.Run()
}