package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgressDb(dbUrl string) *sql.DB {

	if dbUrl == "" {
		log.Fatal("Missind DB_ADDRESS env variable")
	}
	db_conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Couldn't open databse: ", err)
	}
	pingErr := db_conn.Ping()
	if pingErr != nil {
		log.Fatal("Couldn't ping database: ", pingErr)
	}
	return db_conn
}
