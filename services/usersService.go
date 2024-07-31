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