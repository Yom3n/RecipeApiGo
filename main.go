package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	db.NewPostgressDb(dbAddress)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Missing SERVER_PORT env variable")
	}

	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		dat, _ := json.Marshal(map[string]string{"greetings": "hello world"})
		w.Write(dat)
	})
	server := api.NewAPIServer(serverPort, muxHandler)
	log.Fatal(server.Run())
}
