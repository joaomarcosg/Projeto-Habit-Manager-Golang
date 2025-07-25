package utils

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

type ApiResponse struct {
	Error string `json:"error,omitempty"`
	ID    int64  `json:"id,omitempty"`
	Data  any    `json:"data,omitempty"`
}

type HabitRepository interface {
	CreateHabit(ctx context.Context, userID string, habit entity.Habit) (int64, error)
	ListHabits(ctx context.Context, userID string) ([]entity.Habit, error)
	DeleteHabit(ctx context.Context, userID string, id int64) (bool, error)
	UpdateHabit(ctx context.Context, userID string, id int64, habit entity.Habit) (int64, error)
	UpdateHabitStatus(ctx context.Context, userID string, id int64, status string, date time.Time) (bool, error)
	HabitTrack(ctx context.Context, userID string, id int64, startDate, targetDate time.Time) (int64, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

func SendJSON(w http.ResponseWriter, resp ApiResponse, status int) {

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err, "response", resp)
		SendJSON(w, ApiResponse{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}

}
