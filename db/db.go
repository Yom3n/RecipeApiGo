package db

func NewPostgressDb(){

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
}