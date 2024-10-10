package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yom3n/RecipeApiGo/api"
	"github.com/Yom3n/RecipeApiGo/db"
	healthz "github.com/Yom3n/RecipeApiGo/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbAddress := os.Getenv("DB_ADDRESS")
	db.NewPostgressDb(dbAddress)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Missing SERVER_PORT env variable")
	}

	muxHandler := http.NewServeMux()
	
	server := api.NewAPIServer(serverPort, muxHandler)
	log.Fatal(server.Run())
}
