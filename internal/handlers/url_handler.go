package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"url-shortner-app/internal/repository"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/shorten/"):]
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	shortURL, err := repository.CreateShortURL(r.Context(), url)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		log.Printf("CreateShortURL error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortURL)
}

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[len("/fetchOriginalURL/"):]
	if shortCode == "" {
		http.Error(w, "Short code parameter is required", http.StatusBadRequest)
		return
	}

	originalURL, err := repository.FindByShortCode(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "Failed to fetch original URL", http.StatusInternalServerError)
		return
	}
	if originalURL == nil {
		http.Error(w, "Short code not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": originalURL.URL})
}

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[len("/deleteURL/"):]
	if shortCode == "" {
		http.Error(w, "Short code parameter is required", http.StatusBadRequest)
		return
	}

	err := repository.DeleteByShortCode(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "Failed to delete short URL", http.StatusInternalServerError)
		log.Printf("DeleteShortURL error: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
