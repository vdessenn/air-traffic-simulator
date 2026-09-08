package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/engine"
)

type SectorHandler struct {
	sim *engine.Simulation
}

func NewSectorHandler(s *engine.Simulation) *SectorHandler {
	return &SectorHandler{sim: s}
}

func (h *SectorHandler) GetSectors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	sectors := h.sim.GetSectors()
	json.NewEncoder(w).Encode(sectors)
}
