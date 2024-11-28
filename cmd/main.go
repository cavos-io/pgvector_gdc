package main

import (
	"log"
	"net/http"

	"github.com/cavos-io/pgvector_gdc/internal/api"
	"github.com/cavos-io/pgvector_gdc/internal/db"
	"github.com/cavos-io/pgvector_gdc/internal/utils"
)

func main() {
	// Initialize the database connection
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/capabilities", api.CapabilitiesHandler)
	mux.HandleFunc("/schema", api.SchemaHandler)
	mux.HandleFunc("/query", api.QueryHandler)
	mux.HandleFunc("/health", api.HealthHandler)

	// Wrap routes with logging middleware
	loggedRouter := utils.LoggingMiddleware(mux)

	// Start the server on port 8100
	log.Println("Starting server on port 8100...")
	if err := http.ListenAndServe(":8100", loggedRouter); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
