// middleware/cors.go
package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORS(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permita apenas o frontend local
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache pré-checada por 5 minutos
	})(next)
}
