package api

import (
	"net/http"
)

// HealthHandler handles the `/health` endpoint for health checks
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
