package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {

	dsn := "root:root@tcp(localhost:3306)/habits"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	}

	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Printf("connection is not ok: %v", err)
	}

	return db, nil
}
