package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Yom3n/RecipeApiGo/api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Missing SERVER_PORT env variable")
	}
	server := api.NewAPIServer(serverPort)
	log.Fatal(server.Run())
}
