package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	serverAddress := os.Getenv("SERVER_PORT")

	fmt.Println("Starting server at :8080")
	server := &http.Server{
		Addr: serverAddress,
	}
	log.Fatal(server.ListenAndServe())
}
