package object

import "github.com/vdessenn/air-traffic-simulator/backend/internal/model"

type Waypoint struct {
	ID          string           `json:"id"`
	Coordinates model.Coordinate `json:"coordinates"`
}
