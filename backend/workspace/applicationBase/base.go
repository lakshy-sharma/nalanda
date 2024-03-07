package applicationbase

import (
	"fmt"
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"
	"nalanda_backend/middleware"
	"nalanda_backend/routes"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This struct contains the application most commonly shared data.
// It has been created to ensure smooth data sharing between the components of our application.
type NalandaApplication struct {
	appConfig *initializers.AppConfig
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
	DB        *gorm.DB
}

// The entrypoint for the API server.
func ServerEntrypoint() {
	// Read the application config and create a database connection.
	config, err := initializers.LoadConfig(".", "nalanda", "toml")
	if err != nil {
		log.Fatal("Failed to load application settings: ", err)
	}
	dbConn := initializers.ConnectDB(&config)

	// Initialize the application struct and create a GIN server.
	app := &NalandaApplication{
		appConfig: &config,
		InfoLog:   log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:  log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		DB:        dbConn,
	}

	// If operation mode is set to production dont show debug logs.
	if os.Getenv("OPERATION_MODE") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.SetTrustedProxies(nil)

	// This router handles the non authenticated API routes.
	// The middleware we have created is used for handling the CORS.
	routerV1 := server.Group("/api").Group("/v1")
	routerV1.Use(middleware.BaseCorsSettings())

	// Setup Base routes.
	BaseController := controllers.NewBaseController(app.appConfig, app.DB, app.InfoLog, app.ErrorLog)
	BaseRouteController := routes.NewBaseRouteController(BaseController, app.appConfig, app.DB, app.InfoLog, app.ErrorLog)
	BaseRouteController.BaseRoutes(routerV1)

	// Setup authentication routes.
	AuthController := controllers.NewAuthController(app.appConfig, app.DB, app.InfoLog, app.ErrorLog)
	AuthRouteController := routes.NewAuthRouteController(AuthController, app.appConfig, app.DB, app.InfoLog, app.ErrorLog)
	AuthRouteController.AuthRoutes(routerV1)

	// Setup user routes.

	// Start our server and ensure to wrap it inside the error logger to capture any failures.
	app.ErrorLog.Fatal(server.Run(fmt.Sprintf(":%s", app.appConfig.Backend.Port)))
}
