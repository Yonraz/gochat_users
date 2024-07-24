package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonraz/gochat_users/initializers"
	"github.com/yonraz/gochat_users/models"
)

var validSortFields = map[string]bool{
    "status":   true,
    "username": true,
}

var validDirections = map[string]bool{
    "asc":  true,
    "desc": true,
}

func GetUsers(ctx *gin.Context) {
	sort := ctx.DefaultQuery("sort", "status")
	direction := ctx.DefaultQuery("direction", "desc")

	if !validSortFields[sort] {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort field"})
        return
    }

    // Validate sort direction
    if !validDirections[direction] {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort direction"})
        return
    }

	if sort == "online" {
		sort = "status"
	}
	var users []models.User
	if err := initializers.DB.Order(sort + " " + direction).Find(&users).Error; err != nil {
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