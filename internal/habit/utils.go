package habit

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type HabitRepository interface {
	CreateHabit(ctx context.Context, name string) error
}

func sendJSON(w http.ResponseWriter, resp apiResponse, status int) {

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err, "response", resp)
		sendJSON(w, apiResponse{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}

}
