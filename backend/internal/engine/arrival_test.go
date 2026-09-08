package engine

import (
	"testing"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
)

func TestDetectArrivals(t *testing.T) {
	tests := []struct {
		name              string
		setupAircraft     func() map[string]*agent.Aircraft
		expectedCount     int
		expectedCallsigns []string
	}{
		{
			name: "Aircraft with empty flight plan and low speed",
			setupAircraft: func() map[string]*agent.Aircraft {
				aircraft := agent.NewAircraft("TEST001", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{} // Empty flight plan
				aircraft.SetSpeed(30.0)                  // Below landing threshold (50 kts)
				aircraft.PreviousWaypoint = agent.Waypoint{Name: "DEST", Coordinate: agent.Coordinate{}}
				return map[string]*agent.Aircraft{
					"TEST001": aircraft,
				}
			},
			expectedCount:     1,
			expectedCallsigns: []string{"TEST001"},
		},
		{
			name: "Aircraft with empty flight plan but high speed",
			setupAircraft: func() map[string]*agent.Aircraft {
				aircraft := agent.NewAircraft("TEST002", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{}
				aircraft.SetSpeed(150.0) // Above landing threshold
				aircraft.PreviousWaypoint = agent.Waypoint{Name: "DEST"}
				return map[string]*agent.Aircraft{
					"TEST002": aircraft,
				}
			},
			expectedCount:     0,
			expectedCallsigns: []string{},
		},
		{
			name: "Aircraft with waypoints remaining",
			setupAircraft: func() map[string]*agent.Aircraft {
				aircraft := agent.NewAircraft("TEST003", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{
					{Name: "WP1", Coordinate: agent.Coordinate{Latitude: 48.0, Longitude: 2.0}},
				}
				aircraft.SetSpeed(30.0)
				return map[string]*agent.Aircraft{
					"TEST003": aircraft,
				}
			},
			expectedCount:     0,
			expectedCallsigns: []string{},
		},
		{
			name: "Aircraft without previous waypoint (just spawned)",
			setupAircraft: func() map[string]*agent.Aircraft {
				aircraft := agent.NewAircraft("TEST004", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{}
				aircraft.SetSpeed(0.0)
				// No previous waypoint set
				return map[string]*agent.Aircraft{
					"TEST004": aircraft,
				}
			},
			expectedCount:     0,
			expectedCallsigns: []string{},
		},
		{
			name: "Multiple aircraft - mixed states",
			setupAircraft: func() map[string]*agent.Aircraft {
				aircraft1 := agent.NewAircraft("ARRIVED", agent.AircraftType("B737"))
				aircraft1.FlightPlan = []agent.Waypoint{}
				aircraft1.SetSpeed(25.0)
				aircraft1.PreviousWaypoint = agent.Waypoint{Name: "DEST1"}

				aircraft2 := agent.NewAircraft("FLYING", agent.AircraftType("A320"))
				aircraft2.FlightPlan = []agent.Waypoint{
					{Name: "WP1", Coordinate: agent.Coordinate{Latitude: 48.0, Longitude: 2.0}},
				}
				aircraft2.SetSpeed(450.0)

				aircraft3 := agent.NewAircraft("ARRIVED2", agent.AircraftType("B777"))
				aircraft3.FlightPlan = []agent.Waypoint{}
				aircraft3.SetSpeed(0.0)
				aircraft3.PreviousWaypoint = agent.Waypoint{Name: "DEST2"}

				return map[string]*agent.Aircraft{
					"ARRIVED":  aircraft1,
					"FLYING":   aircraft2,
					"ARRIVED2": aircraft3,
				}
			},
			expectedCount:     2,
			expectedCallsigns: []string{"ARRIVED", "ARRIVED2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aircrafts := tt.setupAircraft()
			arrivals := DetectArrivals(aircrafts)

			if len(arrivals) != tt.expectedCount {
				t.Errorf("Expected %d arrivals, got %d", tt.expectedCount, len(arrivals))
			}

			// Check that expected callsigns are present
			arrivalMap := make(map[string]bool)
			for _, callsign := range arrivals {
				arrivalMap[callsign] = true
			}

			for _, expectedCallsign := range tt.expectedCallsigns {
				if !arrivalMap[expectedCallsign] {
					t.Errorf("Expected aircraft %s to be detected as arrived", expectedCallsign)
				}
			}
		})
	}
}

func TestAircraftIsArrived(t *testing.T) {
	tests := []struct {
		name           string
		setupAircraft  func() *agent.Aircraft
		expectedResult bool
	}{
		{
			name: "Arrived aircraft - empty plan and low speed with previous waypoint",
			setupAircraft: func() *agent.Aircraft {
				aircraft := agent.NewAircraft("TEST001", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{}
				aircraft.SetSpeed(20.0)
				aircraft.PreviousWaypoint = agent.Waypoint{Name: "DEST"}
				return aircraft
			},
			expectedResult: true,
		},
		{
			name: "Not arrived - no previous waypoint",
			setupAircraft: func() *agent.Aircraft {
				aircraft := agent.NewAircraft("TEST001B", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{}
				aircraft.SetSpeed(20.0)
				// No previous waypoint
				return aircraft
			},
			expectedResult: false,
		},
		{
			name: "Not arrived - has waypoints",
			setupAircraft: func() *agent.Aircraft {
				aircraft := agent.NewAircraft("TEST002", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{
					{Name: "WP1", Coordinate: agent.Coordinate{Latitude: 48.0, Longitude: 2.0}},
				}
				aircraft.SetSpeed(20.0)
				return aircraft
			},
			expectedResult: false,
		},
		{
			name: "Not arrived - high speed",
			setupAircraft: func() *agent.Aircraft {
				aircraft := agent.NewAircraft("TEST003", agent.AircraftType("B737"))
				aircraft.FlightPlan = []agent.Waypoint{}
				aircraft.SetSpeed(200.0)
				aircraft.PreviousWaypoint = agent.Waypoint{Name: "DEST"}
				return aircraft
			},
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aircraft := tt.setupAircraft()
			result := aircraft.IsArrived()

			if result != tt.expectedResult {
				t.Errorf("Expected IsArrived() to return %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

func TestAircraftShutdown(t *testing.T) {
	aircraft := agent.NewAircraft("TEST001", agent.AircraftType("B737"))

	// Setup initial state
	aircraft.FlightPlan = []agent.Waypoint{
		{Name: "WP1", Coordinate: agent.Coordinate{Latitude: 48.0, Longitude: 2.0}},
	}
	aircraft.SetSpeed(450.0)
	aircraft.Neighbors = map[string]agent.ADSBReport{
		"OTHER": {Emitter: "OTHER"},
	}

	// Call shutdown
	aircraft.Shutdown()

	// Verify cleanup
	if len(aircraft.Neighbors) != 0 {
		t.Errorf("Expected Neighbors to be cleared, got %d entries", len(aircraft.Neighbors))
	}

	if aircraft.FlightPlan != nil {
		t.Errorf("Expected FlightPlan to be nil, got %v", aircraft.FlightPlan)
	}

	if aircraft.GetSpeed() != 0 {
		t.Errorf("Expected Speed to be 0, got %f", aircraft.GetSpeed())
	}
}
