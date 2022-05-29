package config

import (
	"candidate-service/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Instance
var db *gorm.DB

func panicOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s :%s", msg, err)
	}
}

func SetupDatabase() {
	dsn := os.Getenv("DSN")

	// Open connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	panicOnError(err, "Failed to connect database")
	log.Println("Connection open to database")

	// Auto migration
	err = db.AutoMigrate(
		&model.Candidate{},
	)
	panicOnError(err, "Failed to migrate database")
	log.Println("Database migrated")
}
