package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/usecase/habit"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
)

func handleCreateHabit(svc *services.HabitService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := r.Context().Value(userIDKey).(string)
		if !ok || userID == "" {
			utils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "unauthorized",
			})
			return
		}

		data, problems, err := utils.DecodeValidJson[habit.CreateHabitReq](r)
		if err != nil {
			_ = utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
			return
		}

		habit := entity.Habit{
			Name:        data.Name,
			Category:    data.Category,
			Description: data.Description,
			Frequency:   data.Frequency,
			StartDate:   data.StartDate,
			TargetDate:  data.TargetDate,
			Priority:    data.Priority,
		}

		var id int64

		id, err = svc.CreateHabit(r.Context(), userID, habit)
		if err != nil {
			slog.Error("failed to create habit", "error", err)
			utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": "could not create habit",
			})
			return
		}

		utils.EncodeJson(w, r, http.StatusCreated, map[string]any{
			"habit_id": id,
		})

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

func handleUpdateHabit(svc *services.HabitService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := r.Context().Value(userIDKey).(string)
		if !ok || userID == "" {
			utils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "unauthorized",
			})
			return
		}

		data, problems, err := utils.DecodeValidJson[habit.CreateHabitReq](r)
		if err != nil {
			_ = utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
			return
		}

		habit := entity.Habit{
			Name:        data.Name,
			Category:    data.Category,
			Description: data.Description,
			Frequency:   data.Frequency,
			StartDate:   data.StartDate,
			TargetDate:  data.TargetDate,
			Priority:    data.Priority,
		}

		idStr := chi.URLParam(r, "id")
		habitID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
				"error": "invalid param",
			})
			return
		}

		id, err := svc.UpdateHabit(r.Context(), userID, habitID, habit)
		if err != nil {
			slog.Error("failed to update habit", "error", err)
			utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": "failed to update habit",
			})
			return
		}

		utils.EncodeJson(w, r, http.StatusOK, map[string]any{
			"update habit ID": id,
		})

	}
}
