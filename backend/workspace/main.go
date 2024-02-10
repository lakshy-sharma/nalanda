package main

import (
	"fmt"
	"log"
	"nalanda_backend/initializers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type application struct {
	config   initializers.AppConfig
	infoLog  *log.Logger
	errorLog *log.Logger
}

var server *gin.Engine
var app application

// This is the init function which starts our GIN server.
func init() {
	appConfig, err := initializers.LoadConfig(".", "nalanda", "toml")
	if err != nil {
		log.Fatal("Failed to load application settings: ", err)
	}
	initializers.ConnectDB(&appConfig)
	server = gin.Default()
}

// This the main entrypoint for our server and it loads the complete settings.
func main() {
	fmt.Println("Nalanda: An Open Library")

	var err error
	if app.config, err = initializers.LoadConfig(".", "nalanda", "toml"); err != nil {
		app.errorLog.Fatal("Failed to load the environment variables: ", err)
	}

	routerV1 := server.Group("/api").Group("/v1")
	routerV1.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	routerV1.GET("/health", func(ctx *gin.Context) {
		message := "Welcome to Nalanda: The Open Library"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	routerV1.POST("/login", func(ctx *gin.Context) {
		message := "Welcome to Nalanda: The Open Library"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	app.errorLog.Fatal(server.Run(fmt.Sprintf(":%s", app.config.Backend.Port)))
}
