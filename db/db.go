package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) (*gorm.DB, error) {
	dsn := os.Getenv("CONN_STR")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	} else {
		log.Println("database connected successfully")
	}

	return db, nil
}
