package main

import (
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/models"
)

// This function is created to establish a connectio with the database and read the variables quickly.
func init() {
	appConfig, err := initializers.LoadConfig(".", "nalanda", "toml")
	if err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}
	initializers.ConnectDB(&appConfig)
}

// This is why we have a seperate main function for this subsystem.
// The migrations must always be ran apart from the application.
func main() {
	initializers.DB.AutoMigrate(models.User{})
	log.Println("Migrations Completed.")
}
