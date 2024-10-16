package main

import (
	"log"
	"os"

	"github.com/Yom3n/RecipeApiGo/api"
	"github.com/Yom3n/RecipeApiGo/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbAddress := os.Getenv("DB_ADDRESS")
	db := db.NewPostgressDb(dbAddress)
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Missing SERVER_PORT env variable")
	}

	server := api.NewAPIServer(serverPort, db)
	log.Fatal(server.Run())
}
