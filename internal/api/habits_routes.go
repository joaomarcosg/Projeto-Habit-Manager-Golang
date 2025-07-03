package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
)

func NewHabitHandler(svc *services.HabitService) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(AuthMiddleware)

	r.Post("/create", handleCreateHabit(svc))
	r.Get("/list", handleListHabits(svc))
	r.Delete("/delete/{id:[0-9]+}", handleDeleteHabit(svc))

	return r
}
