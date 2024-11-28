package main

import (
	"log"
	"net/http"

	"github.com/cavos-io/pgvector_gdc/internal/api"
	"github.com/cavos-io/pgvector_gdc/internal/db"
)

func main() {
	// Initialize the database connection
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Define routes
	http.HandleFunc("/capabilities", api.CapabilitiesHandler)
	http.HandleFunc("/schema", api.SchemaHandler)
	http.HandleFunc("/query", api.QueryHandler)

	// Start the HTTP server
	log.Println("Starting server on port 8100...")
	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
