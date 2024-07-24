package initializers

import "github.com/yonraz/gochat_users/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}