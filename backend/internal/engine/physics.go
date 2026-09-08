package engine

import (
	"math"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/model"
)

func (s *Simulation) updateAircraftPhysics() {
	for _, plane := range s.Aircrafts {
		speed := 0.000001 * s.SpeedFactor * plane.Speed
		rad := float64(plane.Heading) * math.Pi / 180.0
		dx := speed * math.Cos(rad)
		dy := speed * math.Sin(rad)

		plane.Position.Latitude += dx
		plane.Position.Longitude += dy
	}
}

func (s *Simulation) updateSectorOccupancy() {
	for _, sector := range s.Sectors {
		sector.Counts = 0
		for _, plane := range s.Aircrafts {
			plane.Mu.RLock()
			pos := model.Coordinate{
				Latitude:  plane.Position.Latitude,
				Longitude: plane.Position.Longitude,
			}
			plane.Mu.RUnlock()
			if sector.IsPointInside(pos) {
				sector.Counts++
			}
		}
	}
}

func (s *Simulation) buildConflictAlertLocked(risk model.CollisionRisk) agent.Message {
	var alt1, alt2 int

	if plane, ok := s.Aircrafts[risk.Aircraft1ID]; ok {
		plane.Mu.RLock()
		alt1 = plane.Position.Altitude
		plane.Mu.RUnlock()
	}

	if plane, ok := s.Aircrafts[risk.Aircraft2ID]; ok {
		plane.Mu.RLock()
		alt2 = plane.Position.Altitude
		plane.Mu.RUnlock()
	}

	return agent.Message{
		Type: agent.MessageTypeConflictAlert,
		From: "SIMULATION",
		To:   defaultTowerCallsign,
		Conflict: &agent.ConflictAlert{
			ConflictID:        risk.ID,
			Aircraft1:         risk.Aircraft1ID,
			Aircraft2:         risk.Aircraft2ID,
			Level:             string(risk.Level),
			DistanceMeters:    risk.Distance,
			Aircraft1Altitude: alt1,
			Aircraft2Altitude: alt2,
		},
	}
}

// haversineDistance calculates distance between two coordinates in nautical miles.
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusNM = 3440.065 // Earth radius in nm

	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	lat1Rad := lat1 * (math.Pi / 180.0)
	lat2Rad := lat2 * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusNM * c
}
