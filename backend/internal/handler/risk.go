package handler

import (
	"encoding/json"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/engine"
	"net/http"
)

type RiskHandler struct {
	sim *engine.Simulation
}

func NewRiskHandler(s *engine.Simulation) *RiskHandler {
	return &RiskHandler{sim: s}
}

func (h *RiskHandler) GetRisks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	risks := h.sim.GetRisks()
	json.NewEncoder(w).Encode(risks)
}

func (h *RiskHandler) GetRisksSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	risksSummary := h.sim.GetRisksSummary()
	json.NewEncoder(w).Encode(risksSummary)
}
