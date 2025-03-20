package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	host     = "go_db" // go_db docker -- localhost local
	port     = 5432
)

var dbname, password, user string

func init() {
	// Carrega o arquivo .env
	godotenv.Load()

	dbname = os.Getenv("POSTGRES_DB")
	password = os.Getenv("POSTGRES_PASSWORD")
	user = os.Getenv("POSTGRES_USER")
}

func Connect() (*sql.DB, error) {	
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connected to database %s\n", dbname)
	return db, nil
}