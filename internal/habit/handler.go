package habit

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiResponse struct {
	Error string `json:"error,omitempty"`
	ID    int64  `json:"id,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(svc *Service) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/habits", handleCreateHabit(svc))

	return r
}
