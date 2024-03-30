// This file contains the routes responsible for managing the
package routes

import (
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"
	"nalanda_backend/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminRoutesController struct {
	adminController controllers.AdminController
	appConfig       *initializers.AppConfig
	DB              *gorm.DB
	InfoLogger      *log.Logger
	ErrorLogger     *log.Logger
}

func NewAdminRoutesController(adminController controllers.AdminController, applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) AdminRoutesController {
	return AdminRoutesController{
		adminController: adminController,
		appConfig:       applicationConfig,
		DB:              dbConn,
		InfoLogger:      infoLogger,
		ErrorLogger:     errorLogger,
	}
}

func (adr *AdminRoutesController) AdminRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/admin")
	router.Use(middleware.AuthenticateUser(adr.appConfig, adr.DB, adr.InfoLogger, adr.ErrorLogger))

	router.GET("/all_users", adr.adminController.GetAllUsers)
}
