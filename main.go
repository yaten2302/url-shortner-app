package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"url-shortner-app/internal/db"
	"url-shortner-app/internal/handlers"
	"url-shortner-app/internal/repository"
	"url-shortner-app/internal/router"
	"url-shortner-app/internal/services"
)

func main() {
	// Connect to DB
	DB := db.Connect()
	defer db.Disconnect()

	// Get URLs collection
	collection := DB.Collection("urls")

	urlRepo := repository.NewURLRepository(collection)
	urlService := services.NewURLService(urlRepo)
	urlHandler := handlers.NewURLHandler(urlService)

	// Setup router
	r := router.NewRouter(urlHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start server
	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
