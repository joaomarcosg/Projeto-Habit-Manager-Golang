package api

import (
	"context"
	"net/http"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
)

const userIDKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header missing", http.StatusUnauthorized)
			return
		}

		userID, err := utils.VerifyToken(authHeader)
		if err != nil {
			http.Error(w, "invalid or experied token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
