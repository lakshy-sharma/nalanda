package routes

import (
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRouteController struct {
	userController controllers.UserController
	appConfig      *initializers.AppConfig
	DB             *gorm.DB
	InfoLogger     *log.Logger
	ErrorLogger    *log.Logger
}

func NewUserRouteController(userController controllers.UserController, applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) UserRouteController {
	return UserRouteController{
		userController: userController,
		appConfig:      applicationConfig,
		DB:             dbConn,
		InfoLogger:     infoLogger,
		ErrorLogger:    errorLogger,
	}
}

func (ur *UserRouteController) UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/user")

	router.POST("/me", ur.userController.GetUserData)
}
