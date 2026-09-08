package engine

import (
	"fmt"
	"math"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/model"
)

const (
	SeparationMinDist = 5.0 // 5 NM (Standard)
	CriticalMinDist   = 1.0 // 1 NM (Danger)
)

func DetectConflicts(aircrafts map[string]*agent.Aircraft) []model.CollisionRisk {
	var risks []model.CollisionRisk
	checked := make(map[string]bool)

	for id1, p1 := range aircrafts {
		checked[id1] = true
		for id2, p2 := range aircrafts {
			if checked[id2] {
				continue
			}

			positionP1 := p1.GetPosition()
			positionP2 := p2.GetPosition()

			// Skip detection for aircraft below 100 feet (takeoff/landing phase)
			if positionP1.Altitude < 100 || positionP2.Altitude < 100 {
				continue
			}

			dist := CalculateDistance(positionP1.Latitude, positionP1.Longitude, positionP2.Latitude, positionP2.Longitude)

			altDiff := math.Abs(float64(positionP1.Altitude - positionP2.Altitude))

			// Simplified rule: If less than 5NM horizontally AND less than 1000ft vertically
			if dist < SeparationMinDist && altDiff < 1000 {
				level := model.RiskLevelSafetyNotAssured
				msg := "Minimal separation not assured"

				if dist < CriticalMinDist {
					level = model.RiskLevelCollision
					msg = "IMMEDIATE DANGER: Collision imminent"
				} else if dist < SeparationMinDist/2 {
					level = model.RiskLevelRiskOfCollision
					msg = "High collision risk"
				}

				risks = append(risks, model.CollisionRisk{
					ID:          fmt.Sprintf("%s-%s", id1, id2),
					Aircraft1ID: id1,
					Aircraft2ID: id2,
					Level:       level,
					Distance:    convertNMToMeters(dist),
					Message:     msg,
					Location:    getMiddleCoordinate(positionP1, positionP2),
				})
			}
		}
	}
	return risks
}

func getMiddleCoordinate(c1, c2 agent.Coordinate) model.Coordinate {
	return model.Coordinate{
		Latitude:  (c1.Latitude + c2.Latitude) / 2,
		Longitude: (c1.Longitude + c2.Longitude) / 2,
	}
}

func convertNMToMeters(nm float64) float64 {
	return nm * 1852.0
}

// DetectArrivals checks if any aircraft have reached their final destination.
// An aircraft is considered arrived when:
// 1. It has no more waypoints in its flight plan
// 2. Its speed is below a landing threshold
// 3. It has a previous waypoint (indicating it has been flying a route)
func DetectArrivals(aircrafts map[string]*agent.Aircraft) []string {
	var arrivals []string
	const landingSpeedThreshold = 50.0 // knots

	for callsign, aircraft := range aircrafts {
		aircraft.Mu.RLock()
		flightPlanLen := len(aircraft.FlightPlan)
		speed := aircraft.Speed
		hasPreviousWaypoint := aircraft.PreviousWaypoint.Name != ""
		aircraft.Mu.RUnlock()

		// Aircraft has arrived if:
		// 1. Flight plan is empty (all waypoints reached)
		// 2. Speed is below landing threshold
		// 3. Has a previous waypoint (was following a route)
		if flightPlanLen == 0 && speed < landingSpeedThreshold && hasPreviousWaypoint {
			arrivals = append(arrivals, callsign)
		}
	}

	return arrivals
}
