package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	dsn := os.Getenv("CONN_STR")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	} else {
		log.Println("database connected successfully")
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	fmt.Println("Current time:", now)
}
