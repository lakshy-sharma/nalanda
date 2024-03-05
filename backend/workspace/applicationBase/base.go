package applicationbase

import (
	"fmt"
	"log"
	"nalanda_backend/controllers"
	"nalanda_backend/initializers"
	"nalanda_backend/middleware"
	"nalanda_backend/models"
	"os"

	"github.com/gin-gonic/gin"
)

// This struct contains the application most commonly shared data.
// It has been created to ensure smooth data sharing between the components of our application.
type NalandaApplication struct {
	Config   initializers.AppConfig
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	models   models.Models
}

// The entrypoint for the API server.
func ServerEntrypoint() {
	// Read the application config and create a database connection.
	appConfig, err := initializers.LoadConfig(".", "nalanda", "toml")
	if err != nil {
		log.Fatal("Failed to load application settings: ", err)
	}
	dbConn := initializers.ConnectDB(&appConfig)

	// Initialize the application struct and create a GIN server.
	app := &NalandaApplication{
		Config:   appConfig,
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		models:   models.New(dbConn),
	}

	// If operation mode is set to production dont show debug logs.
	if os.Getenv("OPERATION_MODE") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.SetTrustedProxies(nil)

	// This router handles the non authenticated API routes.
	// The middleware we have created is used for handling the CORS.
	noauthRouterV1 := server.Group("/api").Group("/v1")
	noauthRouterV1.Use(middleware.BaseCorsSettings())
	noauthRouterV1.GET("/", controllers.Index)
	noauthRouterV1.GET("/health", controllers.Healthcheck)
	noauthRouterV1.POST("/login", controllers.Index)

	// Start our server and ensure to wrap it inside the error logger to capture any failures.
	app.ErrorLog.Fatal(server.Run(fmt.Sprintf(":%s", app.Config.Backend.Port)))
}
