package api

import (
	"encoding/json"
	"net/http"
)

// CapabilitiesHandler handles the `/capabilities` endpoint.
func CapabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"version": "1.0",
		"capabilities": map[string]interface{}{
			"queries":   true,
			"mutations": false,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SchemaHandler handles the `/schema` endpoint.
func SchemaHandler(w http.ResponseWriter, r *http.Request) {
	// Return static schema for now
	w.WriteHeader(http.StatusOK)
}

// QueryHandler handles the `/query` endpoint.
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
