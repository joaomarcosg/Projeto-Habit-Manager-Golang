package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/services"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/store/mysqlstore"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/usecase/user"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
)

func handleSignupUser(svc *services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, problems, err := utils.DecodeValidJson[user.CreateUserReq](r)
		if err != nil {
			_ = utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
			return
		}
		id, err := svc.CreateUser(r.Context(), data.Name, data.Email, data.Password)
		if err != nil {
			if errors.Is(err, mysqlstore.ErrDuplicatedEmailOrUsername) {
				slog.Error("failed to create user", "error", err)
				_ = utils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
					"error": "email or username already exist",
				})
				return
			}
		}
		_ = utils.EncodeJson(w, r, http.StatusCreated, map[string]any{
			"user_id": id.ID,
		})
	}

}
