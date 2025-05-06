package mysqlstore

import (
	"database/sql"
	"log"
)

func ConnectDB(datasource string) (*sql.DB, error) {

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Printf("Error connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping MySQL: %v", err)
	}

	return db, nil
}
