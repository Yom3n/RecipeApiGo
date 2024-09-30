package main

import (
	"github.com/Yom3n/RecipeApiGo/api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	// serverAddress := os.Getenv("SERVER_PORT")
	server := api.NewAPIServer(":8080")
	log.Fatal(server.Run())
}
