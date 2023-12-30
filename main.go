package main

import (
	"log"

	"github.com/Darshanbennur/Trade-Book/cmd"
	"github.com/Darshanbennur/Trade-Book/db"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load environment vairables", err)
	}

	_ = db.Start()

	cmd.Execute()

}
