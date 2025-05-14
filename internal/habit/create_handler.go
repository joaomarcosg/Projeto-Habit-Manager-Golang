package habit

import (
	"encoding/json"
	"net/http"
)

type inputHabit struct {
	Name string `json:"name"`
}

func handleCreateHabit(svc *Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var input inputHabit

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			sendJSON(w, apiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if err := svc.CreateHabit(r.Context(), input.Name); err != nil {
			sendJSON(w, apiResponse{Error: "could not create habit"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: input}, http.StatusCreated)
	}
}
