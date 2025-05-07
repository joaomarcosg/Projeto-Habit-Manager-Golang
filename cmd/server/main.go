package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/habit"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/store/mysqlstore"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("all system offline")

}

func run() error {

	dsn := "mysql://root:root@tcp(localhost:3306)/habits"
	sqlDB, err := mysqlstore.ConnectDB(dsn)
	if err != nil {
		log.Printf("failed to connect: %v", err)
	}

	handler := habit.NewHandler(sqlDB)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
