package services

import (
	"fmt"

	"github.com/yonraz/gochat_users/constants"
	"github.com/yonraz/gochat_users/initializers"
	"github.com/yonraz/gochat_users/models"
	"github.com/yonraz/gochat_users/state"
)


func Create(user models.User) error {
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("could not create a new user: %v", result.Error)
	}
	fmt.Printf("saved user %v to database!\n", user.Username)
	state.DbCacheState.SetIsChanged(true)
	return nil
}

func Login(user models.User) error {
	err := initializers.DB.Model(&models.User{}).Where("username = ?", user.Username).Update("Status", constants.Online).Error
	if err != nil {
		fmt.Printf("error updating user: %v\n", err)
		return err
	}
	state.DbCacheState.SetIsChanged(true)
	return nil
}

func Signout(user models.User) error {
	err := initializers.DB.Model(&models.User{}).Where("username = ?", user.Username).Update("Status", constants.Offline).Error
	if err != nil {
		fmt.Printf("error updating user: %v\n", err)
		return err
	}
	state.DbCacheState.SetIsChanged(true)
	return nil
}

func CreateMockUsers() {
	users := []models.User{
        {Username: "alice", Status: constants.Online},
        {Username: "charlie", Status: constants.Online},
        {Username: "eve", Status: constants.Online},
        {Username: "bob", Status: constants.Offline},
        {Username: "grace", Status: constants.Online},
        {Username: "dave", Status: constants.Offline},
        {Username: "frank", Status: constants.Offline},
        {Username: "heidi", Status: constants.Offline},
        {Username: "ivan", Status: constants.Online},
        {Username: "judy", Status: constants.Offline},
        {Username: "Dan", Status: constants.Offline},
        {Username: "daniel", Status: constants.Offline},
        {Username: "Ben", Status: constants.Offline},
        {Username: "Ron", Status: constants.Offline},
        {Username: "john", Status: constants.Online},
        {Username: "jane", Status: constants.Offline},
        {Username: "mike", Status: constants.Online},
        {Username: "lisa", Status: constants.Offline},
        {Username: "paul", Status: constants.Online},
        {Username: "anna", Status: constants.Offline},
        {Username: "tina", Status: constants.Online},
        {Username: "mark", Status: constants.Offline},
        {Username: "susan", Status: constants.Online},
        {Username: "tom", Status: constants.Offline},
        {Username: "lucas", Status: constants.Online},
        {Username: "maria", Status: constants.Offline},
        {Username: "oliver", Status: constants.Online},
        {Username: "claire", Status: constants.Offline},
        {Username: "ryan", Status: constants.Online},
        {Username: "kate", Status: constants.Offline},
        {Username: "nate", Status: constants.Online},
        {Username: "laura", Status: constants.Offline},
        {Username: "sophie", Status: constants.Online},
        {Username: "eric", Status: constants.Offline},
        {Username: "jessica", Status: constants.Online},
        {Username: "george", Status: constants.Offline},
        {Username: "hannah", Status: constants.Online},
        {Username: "steve", Status: constants.Offline},
        {Username: "lily", Status: constants.Online},
        {Username: "leo", Status: constants.Offline},
    }

    // Insert the users into the database
    result := initializers.DB.Create(&users)
    if result.Error != nil {
        panic("Failed to insert records")
    }

    fmt.Println("Records inserted successfully")
}