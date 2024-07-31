package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yonraz/gochat_users/cache"
	"github.com/yonraz/gochat_users/initializers"
	"github.com/yonraz/gochat_users/models"
	"github.com/yonraz/gochat_users/state"
)

var validSortFields = map[string]bool{
    "status":   true,
    "username": true,
}

var validDirections = map[string]bool{
    "asc":  true,
    "desc": true,
}

var PAGE_SIZE = 20

type UsersController struct {
	cacheInstance *cache.Redis
}

func NewUsersController() *UsersController {
	cacheInstance, err := cache.NewClient()
	if err != nil {
		log.Printf("error creating cache instance")
	}
	return &UsersController{
		cacheInstance: cacheInstance,
	}
}

func (controller *UsersController) GetUsers(ctx *gin.Context) {
	sort := ctx.DefaultQuery("sort", "status")
	direction := ctx.DefaultQuery("direction", "desc")
	pageQuery := ctx.DefaultQuery("page", "1")
	page, queryErr := strconv.Atoi(pageQuery)
	if queryErr != nil {
		log.Println("page query invalid, defaulting to 1.")
		pageQuery = "1"
		page = 1
	}
	OFFSET := (page - 1) * PAGE_SIZE

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

	query := sort + direction + "page:" + pageQuery

	var users []models.User
	if state.DbCacheState.WasDBChanged() {
		result, err := controller.cacheInstance.GetQuery(query)
		if err == nil && result != nil {
			ctx.JSON(http.StatusOK, gin.H{
			"users": result,
			})
			log.Println(query + " query hit")
		} else {
			log.Println(query + " query miss")
		}
	}

	if err := initializers.DB.Offset(OFFSET).Limit(PAGE_SIZE).Order(sort + " " + direction).Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
            "details": err.Error(),
		})
		return
	}
	err := controller.cacheInstance.SetQuery(query, users)
	state.DbCacheState.SetIsChanged(false)
	if err != nil {
		log.Printf("error setting redis cache, %v\n", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}