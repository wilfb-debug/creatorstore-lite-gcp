package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
	DB *sql.DB
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {

	err := h.DB.Ping()

	response := map[string]string{
		"status": "ok",
	}

	if err != nil {
		response["database"] = "disconnected"
	} else {
		response["database"] = "connected"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
