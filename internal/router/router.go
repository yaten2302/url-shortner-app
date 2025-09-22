package router

import (
	"url-shortner-app/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(urlHandler *handlers.URLHandler) *chi.Mux {
	r := chi.NewRouter()

	// Routes
	r.Post("/shorten", urlHandler.CreateShortURL)
	r.Get("/fetchOriginalURL", urlHandler.GetOriginalURL)

	return r
}
