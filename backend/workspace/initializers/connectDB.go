package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(applicationConfig *AppConfig) *gorm.DB {
	var err error

	// Database connection string.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", applicationConfig.Database.Host, applicationConfig.Database.User, applicationConfig.Database.Password, applicationConfig.Database.DatabaseName, applicationConfig.Database.Port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect with the database.")
	}
	log.Println("Established connection with the database.")
	return DB
}
