package controllers

import (
	"log"
	"nalanda_backend/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
	appConfig   *initializers.AppConfig
	DB          *gorm.DB
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewBaseController(applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) BaseController {
	return BaseController{
		appConfig:   applicationConfig,
		DB:          dbConn,
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}
}

func (bc *BaseController) Healthcheck(ctx *gin.Context) {
	message := "Welcome to Nalanda: Nalanda: An Open Knowledge Center"
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (bc *BaseController) Index(ctx *gin.Context) {
	message := "Welcome to Nalanda: Nalanda: An Open Knowledge Center"
	ctx.JSON(http.StatusOK, gin.H{"message": message, "status": "success"})
}
