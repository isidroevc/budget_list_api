package supabase

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db_singleton *sql.DB

func GetConnection() (*sql.DB, error) {

	if db_singleton != nil {
		return db_singleton, nil
	}

	connStr := os.Getenv("SUPA_BASE_CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)
	db.SetMaxOpenConns(10)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to initialize supabase external service due to error %s", err.Error())
	}
	db_singleton = db
	return db_singleton, nil

}
