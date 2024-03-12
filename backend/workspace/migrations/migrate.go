package main

import (
	"log"
	"nalanda_backend/initializers"
	appModels "nalanda_backend/models"
)

// This is why we have a seperate main function for this subsystem.
// The migrations must always be ran apart from the application.
func main() {
	// Read configurations and connect to the database.
	appConfig, err := initializers.LoadConfig(".", "nalanda", "toml")
	if err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}
	db := initializers.ConnectDB(&appConfig)
	// Perform the migrations to the database.
	db.AutoMigrate(appModels.User{})
	log.Println("Migrations Completed.")
}
