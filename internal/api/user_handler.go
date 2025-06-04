package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleSignupUser(svc *UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user CreateUserReq

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendJSON(w, apiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		id, err := svc.CreateUser(r.Context(), user.Name, user.Email, user.Password)
		if err != nil {
			slog.Error("failed to create user", "error", err)
			sendJSON(w, apiResponse{Error: "could not create user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: id.ID}, http.StatusCreated)

	}

}
