package handler

import (
	"encoding/json"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/engine"
	"net/http"
)

type PlaneHandler struct {
	sim *engine.Simulation
}

func NewPlaneHandler(s *engine.Simulation) *PlaneHandler {
	return &PlaneHandler{sim: s}
}

func (h *PlaneHandler) GetPlanes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	planes := h.sim.GetAircrafts()
	json.NewEncoder(w).Encode(planes)
}
