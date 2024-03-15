package controllers

import (
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	appConfig   *initializers.AppConfig
	DB          *gorm.DB
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewUserController(applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) UserController {
	return UserController{
		appConfig:   applicationConfig,
		DB:          dbConn,
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}
}

// This function is used for retrieving users by ID.
func (uc *UserController) GetUserData(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	userResponse := &models.UserDataResponse{
		ID:        currentUser.ID,
		Email:     currentUser.Email,
		Role:      currentUser.Role,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
