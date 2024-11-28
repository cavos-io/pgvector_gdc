package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cavos-io/pgvector_gdc/internal/db"
	"github.com/pgvector/pgvector-go"
)

type QueryRequest struct {
	Vector []float32 `json:"vector"`
	TopK   int       `json:"top_k"`
}

type QueryResponse struct {
	Results []map[string]interface{} `json:"results"`
}

// QueryHandler handles the `/query` endpoint.
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	var queryReq QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&queryReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Convert input vector to pgvector.Vector
	vector := pgvector.NewVector(queryReq.Vector)

	// Query the nearest neighbors
	results, err := db.FindNearestVectors(vector, queryReq.TopK)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error querying vectors: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Return results
	response := QueryResponse{Results: results}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
