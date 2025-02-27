package helpers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("your_secret_key")

type Claims struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(email, userID string) (string, error) {
	claims := &Claims{
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "auth_service",
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
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

		claims, err := ValidateToken(tokenParts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		r = mux.SetURLVars(r, map[string]string{"user_id": claims.UserID})
		next.ServeHTTP(w, r)
	})
}
