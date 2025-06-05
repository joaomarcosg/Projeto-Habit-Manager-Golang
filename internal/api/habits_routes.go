package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
)

func NewHandler(svc *services.HabitService) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/habits", handleCreateHabit(svc))
	r.Get("/habits/list", handleListHabits(svc))
	r.Delete("/habits/delete/{id:[0-9]+}", handleDeleteHabit(svc))

	return r
}
