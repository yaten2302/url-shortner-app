package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"url-shortner-app/internal/services"
)

type URLHandler struct {
	service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.CreateShortURL(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		log.Printf("CreateShortURL error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortURL)
}

func (h *URLHandler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ShortCode string `json:"short_code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ShortCode == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	originalURL, err := h.service.GetOriginalURL(r.Context(), req.ShortCode)
	if err != nil {
		http.Error(w, "Failed to fetch original URL", http.StatusInternalServerError)
		return
	}
	if originalURL == "" {
		http.Error(w, "Short code not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": originalURL})
}

func (h *URLHandler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ShortCode string `json:"short_code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ShortCode == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteURL(r.Context(), req.ShortCode)
	if err != nil {
		http.Error(w, "Failed to delete short URL", http.StatusInternalServerError)
		log.Printf("DeleteShortURL error: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
