package middleware

import (
	"net/http"
)

const APIKey = "qwerklj1230dsa350123l2k1j4kl1j24"

func APIKeyValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("APIKey")
		if apiKey != APIKey {
			http.Error(w, "Forbidden: Invalid APIKey", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
