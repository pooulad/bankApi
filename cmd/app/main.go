package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pooulad/bankApi/api"
)

func main() {

	config, err := config.ReadPostgresConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	host := os.Getenv("PROJECT_HOST")
	port := os.Getenv("PROJECT_PORT")

	serverAddress := fmt.Sprintf("%s:%s", host, port)

	server := api.NewApiServer(serverAddress)
	server.Run()
}
