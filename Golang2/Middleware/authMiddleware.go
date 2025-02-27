package middleware

import (
	"errors"
	"net/http"
	"strings"

	helper "github.com/devesh/mongoapi/Helpers"
	"github.com/gorilla/mux"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid auth token format", http.StatusUnauthorized)
			return
		}

		claims, err := helper.ValidateToken(tokenParts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		r = mux.SetURLVars(r, map[string]string{"user_id": claims.UserID})
		next.ServeHTTP(w, r)
	})
}