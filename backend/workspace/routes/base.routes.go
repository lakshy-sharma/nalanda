package routes

import (
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseRouteController struct {
	BaseController controllers.BaseController
}

func NewBaseRouteController(baseController controllers.BaseController, applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) BaseRouteController {
	return BaseRouteController{baseController}
}

func (br *BaseRouteController) BaseRoutes(rg *gin.RouterGroup) {
	router := rg
	router.GET("/", br.BaseController.Index)
	router.GET("/health", br.BaseController.Healthcheck)
}
