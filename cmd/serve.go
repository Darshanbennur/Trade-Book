package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func Execute() {
	server = gin.Default()
	fmt.Println("Initializing the server ...")
	log.Fatal(server.Run(":8080"))
}
