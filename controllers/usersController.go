package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonraz/gochat_users/initializers"
	"github.com/yonraz/gochat_users/models"
)

func GetUsers(ctx *gin.Context) {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
            "details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}