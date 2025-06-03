package habit

import (
	"encoding/json"
	"net/http"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleSignupUser(svc *Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user CreateUserReq

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendJSON(w, apiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

	}

}
