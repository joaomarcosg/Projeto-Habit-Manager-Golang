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

func handleLoginUser(svc *services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data, problems, err := utils.DecodeValidJson[user.LoginUserReq](r)
		if err != nil {
			utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		}
		id, token, err := svc.AuthenticateUser(r.Context(), data.Email, data.Password)
		if err != nil {
			if errors.Is(err, mysqlstore.ErrInvalidCredentials) {
				utils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
					"error": "invalid email or password",
				})
				return
			}
			utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": "unexpected server error",
			})
			return
		}

		utils.EncodeJson(w, r, http.StatusOK, map[string]any{
			"user_id": id.ID,
			"token":   token,
		})

	}

}
