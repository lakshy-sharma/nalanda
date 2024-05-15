package routes

import (
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"
	"nalanda_backend/middleware"

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
	router.Use(middleware.AuthenticateUser(ur.appConfig, ur.DB, ur.InfoLogger, ur.ErrorLogger))

	router.GET("/me", ur.userController.GetUserData)
	router.DELETE("/me", ur.userController.DeleteUser)
	router.POST("/update", ur.userController.UpdateUser)
}
