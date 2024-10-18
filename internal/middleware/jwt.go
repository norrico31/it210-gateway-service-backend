package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const BEARER = "Bearer "

func ValidateJWT(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, BEARER) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BEARER)
		jwtSecret, exists := os.LookupEnv("JWT_SECRET")
		if !exists {
			http.Error(w, "Gateway: env variables not set", http.StatusInternalServerError)
			return
		}
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		n.ServeHTTP(w, r)
	})
}
