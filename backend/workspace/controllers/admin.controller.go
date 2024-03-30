package controllers

import (
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	appConfig   *initializers.AppConfig
	DB          *gorm.DB
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewAdminController(applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) AdminController {
	return AdminController{
		appConfig:   applicationConfig,
		DB:          dbConn,
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}
}

func (adc *AdminController) GetAllUsers(ctx *gin.Context) {
	var users []models.User
	// Fetch a list of all users using GORM.
	result := adc.DB.Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Failed to fetch all users."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": users})
}
