package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"url-shortner-app/internal/db"
	"url-shortner-app/internal/router"
)

func main() {
	// Connect to DB
	_ = db.Connect()
	defer db.Disconnect()

	// Setup router
	r := router.NewRouter()

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
