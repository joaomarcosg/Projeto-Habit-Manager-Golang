package mysqlstore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(datasource string) (*sql.DB, error) {

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Printf("error connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("failed to ping MySQL: %v", err)
	}

	return db, nil
}
