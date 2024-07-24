package controllers

import (
	"fmt"

	"github.com/yonraz/gochat_users/initializers"
	"github.com/yonraz/gochat_users/models"
)

func Signup(username string) error {
	user := models.User{Username: username}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("could not create a new user: %v", result.Error)
	}
	fmt.Printf("saved user %v to database!", username)
	return nil
}