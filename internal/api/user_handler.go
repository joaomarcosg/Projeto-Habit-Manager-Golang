package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/service"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleSignupUser(svc *service.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user CreateUserReq

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.SendJSON(w, utils.ApiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		id, err := svc.CreateUser(r.Context(), user.Name, user.Email, user.Password)
		if err != nil {
			slog.Error("failed to create user", "error", err)
			utils.SendJSON(w, utils.ApiResponse{Error: "could not create user"}, http.StatusInternalServerError)
			return
		}

		utils.SendJSON(w, utils.ApiResponse{Data: id.ID}, http.StatusCreated)

	}

}
