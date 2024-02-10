package applicationbase

import (
	"fmt"
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This struct contains the application most commonly shared data.
// It has been created to ensure smooth data sharing between the components of our application.
type NalandaApplication struct {
	Config   initializers.AppConfig
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Db       *gorm.DB
}

// The entrypoint for the API server.
func ServerEntrypoint() {
	// Read the application configurations and create a database connection.
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
		Db:       dbConn,
	}

	// If operation mode is set to production dont show debug logs.
	if os.Getenv("OPERATION_MODE") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.SetTrustedProxies(nil)

	// Create a router to handle the various API routes.
	routerV1 := server.Group("/api").Group("/v1")
	routerV1.Use(middleware.BaseCorsSettings())

	// Some basic handlers.
	// TODO move them into controllers.
	routerV1.GET("/health", func(ctx *gin.Context) {
		message := "Welcome to Nalanda: The Open Library"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	routerV1.POST("/login", func(ctx *gin.Context) {
		message := "Welcome to Nalanda: The Open Library"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// Start our server and ensure to wrap it inside the error logger to capture any failures.
	app.ErrorLog.Fatal(server.Run(fmt.Sprintf(":%s", app.Config.Backend.Port)))
}
