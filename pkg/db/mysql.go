package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {

	db, err := sql.Open("mysql", "root:root@habits")
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	}

	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Printf("connection is not ok: %v", err)
	}

}
