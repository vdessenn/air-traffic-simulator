package handler

import (
	"encoding/json"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/engine"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SettingsHandler struct {
	sim *engine.Simulation
}

func NewSettingsHandler(s *engine.Simulation) *SettingsHandler {
	return &SettingsHandler{sim: s}
}

func (h *SettingsHandler) SetSpeedFactor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	valueStr := strings.TrimSpace(string(body))

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Body must be a float64", http.StatusBadRequest)
		return
	}

	h.sim.SpeedFactor = value
	log.Printf("Changement de vitesse d'exécution : %s", body)

	json.NewEncoder(w).Encode("OK")
}
