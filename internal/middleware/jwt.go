package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const BEARER = "Bearer "

// ValidateJWT validates the JWT token and extracts the userId, adding it to the headers.
func ValidateJWT(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, BEARER) {
			http.Error(w, "Authorization header missing or invalid", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BEARER)
		jwtSecret, exists := os.LookupEnv("JWT_SECRET")
		if !exists {
			http.Error(w, "Gateway: env variable JWT_SECRET not set", http.StatusInternalServerError)
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims and validate expiration
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				http.Error(w, "Token has expired", http.StatusUnauthorized)
				return
			}
		}

		// Extract userId from claims
		userID, ok := claims["userId"].(string) // Adjust this if `userId` is stored as an integer
		if !ok || userID == "" {
			http.Error(w, "userId not found in token", http.StatusUnauthorized)
			return
		}

		// Add userId to the headers for downstream use
		r.Header.Set("X-User-ID", userID)

		// Proceed with the request
		h.ServeHTTP(w, r)
	})
}
