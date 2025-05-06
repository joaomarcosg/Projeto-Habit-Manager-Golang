package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/store/mysqlstore"
)

func main() {

}

func run() error {

	dsn := "mysql://root:root@tcp(localhost:3306)/habits"
	sqlDB, err := mysqlstore.ConnectDB(dsn)
	if err != nil {
		log.Printf("Failed to connect: %v", err)
	}

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
