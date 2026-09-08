package handler

import (
	"encoding/json"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/engine"
	"net/http"
)

type WaypointHandler struct {
	sim *engine.Simulation
}

func NewWaypointHandler(s *engine.Simulation) *WaypointHandler {
	return &WaypointHandler{sim: s}
}

func (h *WaypointHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	tags := h.sim.GetWaypoints()
	json.NewEncoder(w).Encode(tags)
}
