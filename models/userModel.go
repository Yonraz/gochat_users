package models

import (
	"github.com/yonraz/gochat_users/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Status constants.UserStatus
}