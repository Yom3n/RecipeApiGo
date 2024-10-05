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
	dbAddress := os.Getenv("DB_ADDRESS")
	if dbAddress == "" {
		log.Fatal("Missind DB_ADDRESS env variable")
	}
	db_conn, err := sql.Open("postgres", dbAddress)
	if err != nil {
		log.Fatal("Couldn't open databse: ", err)
	}
	pingErr := db_conn.Ping()
	if pingErr != nil {
		log.Fatal("Couldn't ping database: ", pingErr)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Missing SERVER_PORT env variable")
	}
	server := api.NewAPIServer(serverPort)
	log.Fatal(server.Run())
}
