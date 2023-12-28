package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Darshanbennur/Trade-Book/api/handlers"
	"github.com/Darshanbennur/Trade-Book/api/services"
	"github.com/Darshanbennur/Trade-Book/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	server               *gin.Engine
	database             gorm.DB
	ctx                  context.Context
	instrumentService    services.InstrumentServices
	tradeService         services.TradeServices
	instrumentController handlers.InstrumentController
	tradeController      handlers.TradeController
)

func init() {
	ctx = context.TODO()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load environment vairables", err)
	}

	database, error := db.Start(&database)
	if error != nil {
		// Handle error
	}

	instrumentService = services.NewInstrumentService(database, ctx)
	tradeService = services.NewTradeService(database, ctx)
	instrumentController = handlers.NewInstrument(instrumentService)
	tradeController = handlers.NewTrade(tradeService)

	server = gin.Default()
}

func main() {

	basePath := server.Group("/v1")
	tradeController.RegisterTradeRoutes(basePath)
	instrumentController.RegisterInstrumentRoutes(basePath)

	//Similarly rest Basepath and their controllers will be written here

	fmt.Println("Initializing the server ...")
	log.Fatal(server.Run(":8080"))
}
