package habit

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

type inputHabit struct {
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Frequency   string    `json:"frequency"`
	StartDate   time.Time `json:"start_date"`
	TargetDate  time.Time `json:"target_date"`
	Priority    uint8     `json:"priority"`
}

func handleCreateHabit(svc *Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var input inputHabit

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			sendJSON(w, apiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		habit := entity.Habit{
			Name:        input.Name,
			Category:    input.Category,
			Description: input.Description,
			Frequency:   input.Frequency,
			StartDate:   input.StartDate,
			TargetDate:  input.TargetDate,
			Priority:    input.Priority,
		}

		var id int64

		id, err := svc.CreateHabit(r.Context(), habit)
		if err != nil {
			slog.Error("failed to create habit", "error", err)
			sendJSON(w, apiResponse{Error: "could not create habit"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{ID: id}, http.StatusCreated)
	}
}

func handleListHabits(svc *Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		habits, err := svc.ListHabits(r.Context())
		if err != nil {
			slog.Error("failed to get habits", "error", err)
			sendJSON(w, apiResponse{Error: "could not list habits"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: habits}, http.StatusOK)

	}

}

func handleDeleteHabit(svc *Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			sendJSON(w, apiResponse{Error: "invalid param"}, http.StatusBadRequest)
			return
		}

		ok, err := svc.DeleteHabit(r.Context(), id)

		if err != nil {
			slog.Error("failed to delete habit", "error", err)
			sendJSON(w, apiResponse{Error: "failed to delete habit"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: ok}, http.StatusOK)

	}
}
