package main

import (
	"backend/config"
	"backend/rooter"

	"log"
)

func main() {
	config.InitConfig()

	r := rooter.SetupRooter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Server startup failure: %v", err)
	}
}
