package router

import (
	"url-shortner-app/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	// Routes
	r.Post("/shorten/{url}", handlers.CreateShortURL)
	r.Get("/fetchOriginalURL/{short_code}", handlers.GetOriginalURL)
	r.Delete("/deleteURL/{short_code}", handlers.DeleteURL)

	return r
}
