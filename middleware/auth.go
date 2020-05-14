package middleware

import (
	"net/http"
	"os"
)

// AuthMiddleware function that checks if the apiKey is correct for operation
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		envAPIKey := os.Getenv("API_KEY")
		apiKey := r.Header.Get("apiKey")

		if apiKey != envAPIKey {
			http.Error(w, "You are not authorized to access this page", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})

}
