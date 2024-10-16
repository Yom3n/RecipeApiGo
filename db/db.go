package db

import (
	"database/sql"
	"log"

	"github.com/Yom3n/RecipeApiGo/db/database"
	_ "github.com/lib/pq"
)

func NewPostgressDb(dbUrl string) *database.Queries {

	if dbUrl == "" {
		log.Fatal("Missind DB_ADDRESS env variable")
	}
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Couldn't open databse: ", err)
	}
	pingErr := dbConn.Ping()
	if pingErr != nil {
		log.Fatal("Couldn't ping database: ", pingErr)
	}
	log.Print("Connected to database")
	return database.New(dbConn)
}
