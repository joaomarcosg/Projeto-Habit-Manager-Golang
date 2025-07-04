package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
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

func handleCreateHabit(svc *services.HabitService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var input inputHabit

		userID, ok := r.Context().Value(userIDKey).(string)
		if !ok || userID == "" {
			utils.SendJSON(w, utils.ApiResponse{Error: "unauthorized"}, http.StatusUnauthorized)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.SendJSON(w, utils.ApiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
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

		id, err := svc.CreateHabit(r.Context(), userID, habit)
		if err != nil {
			slog.Error("failed to create habit", "error", err)
			utils.SendJSON(w, utils.ApiResponse{Error: "could not create habit"}, http.StatusInternalServerError)
			return
		}

		utils.SendJSON(w, utils.ApiResponse{ID: id}, http.StatusCreated)
	}
}

func handleListHabits(svc *services.HabitService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := r.Context().Value(userIDKey).(string)
		if !ok || userID == "" {
			utils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "unauthorized",
			})
			return
		}

		habits, err := svc.ListHabits(r.Context(), userID)
		if err != nil {
			slog.Error("failed to get habits", "error", err)
			utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": "could not list habits",
			})
			return
		}

		utils.EncodeJson(w, r, http.StatusOK, map[string]any{
			"habtis": habits,
		})

	}

}

func handleDeleteHabit(svc *services.HabitService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := r.Context().Value(userIDKey).(string)
		if !ok || userID == "" {
			utils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "unauthorized",
			})
			return
		}

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
				"error": "invalid param",
			})
			return
		}

		ok, err = svc.DeleteHabit(r.Context(), userID, id)
		if err != nil {
			slog.Error("failed to delete habit", "error", err)
			utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": "failed to delete habit",
			})
			return
		}

		utils.EncodeJson(w, r, http.StatusOK, map[string]any{
			"deleted": ok,
		})

	}
}
