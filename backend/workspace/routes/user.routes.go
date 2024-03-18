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

	router.GET("/me", middleware.AuthenticateUser(ur.appConfig, ur.DB, ur.InfoLogger, ur.ErrorLogger), ur.userController.GetUserData)
	router.POST("/update_email", middleware.AuthenticateUser(ur.appConfig, ur.DB, ur.InfoLogger, ur.ErrorLogger), ur.userController.UpdateUserEmail)
	router.POST("/reset_password", middleware.AuthenticateUser(ur.appConfig, ur.DB, ur.InfoLogger, ur.ErrorLogger), ur.userController.ResetUserPassword)
	router.DELETE("/me", middleware.AuthenticateUser(ur.appConfig, ur.DB, ur.InfoLogger, ur.ErrorLogger), ur.userController.DeleteUser)
}
