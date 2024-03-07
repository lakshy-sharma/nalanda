package routes

import (
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"
	"nalanda_backend/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRouteController struct {
	authController controllers.AuthController
	appConfig      *initializers.AppConfig
	DB             *gorm.DB
	InfoLogger     *log.Logger
	ErrorLogger    *log.Logger
}

func NewAuthRouteController(authController controllers.AuthController, applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) AuthRouteController {
	return AuthRouteController{
		authController,
		applicationConfig,
		dbConn,
		infoLogger,
		errorLogger,
	}
}

func (ar *AuthRouteController) AuthRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", ar.authController.UserSignUp)
	router.POST("/login", ar.authController.UserSignIn)
	router.GET("/logout", middleware.DeserializeUser(ar.appConfig, ar.DB, ar.InfoLogger, ar.ErrorLogger), ar.authController.LogoutUser)
}
