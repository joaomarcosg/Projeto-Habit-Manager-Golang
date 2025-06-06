package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/api"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
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

	dsn := "root:root@tcp(localhost:3306)/habits?parseTime=true"
	sqlDB, err := mysqlstore.ConnectDB(dsn)
	if err != nil {
		log.Printf("failed to connect: %v", err)
	}

	queries := mysqlstore.New(sqlDB)

	habitRepo := mysqlstore.NewHabitRepository(queries)
	habitSvc := services.NewHabitService(habitRepo)

	userRepo := mysqlstore.NewUserRepository(queries)
	userSvc := services.NewUserService(userRepo)

	r := chi.NewRouter()

	r.Mount("/habits", api.NewHabitHandler(habitSvc))
	r.Mount("/users", api.NewUserHandler(userSvc))

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      r,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
