package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
)

func NewUserHandler(svc *services.UserService) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/signup", handleSignupUser(svc))

	return r

}
