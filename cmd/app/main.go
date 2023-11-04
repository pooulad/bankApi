package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pooulad/bankApi/api"
	"github.com/pooulad/bankApi/config"
	"github.com/pooulad/bankApi/database"
)

func main() {

	config, err := config.ReadPostgresConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	storage, err := database.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", storage)

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	host := os.Getenv("PROJECT_HOST")
	port := os.Getenv("PROJECT_PORT")

	serverAddress := fmt.Sprintf("%s:%s", host, port)

	server := api.NewApiServer(serverAddress,storage)
	server.Run()
}
