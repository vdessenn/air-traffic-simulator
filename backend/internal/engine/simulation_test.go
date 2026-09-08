package engine

import (
	"testing"
	"time"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
)

// waitFor condition helper with timeout.
func waitFor(t *testing.T, timeout time.Duration, cond func() bool) bool {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if cond() {
			return true
		}
		time.Sleep(10 * time.Millisecond)
	}
	return false
}

func TestConflictTriggersDirectiveAndAltitudeChange(t *testing.T) {
	sim := NewSimulation()
	defer sim.Stop()

	// Two aircraft on near-collision course at same altitude.
	p1 := agent.NewAircraft("AFR100", "COMMERCIAL")
	p1.Position = agent.Coordinate{Latitude: 0, Longitude: 0, Altitude: 30000}

	p2 := agent.NewAircraft("AFR200", "COMMERCIAL")
	p2.Position = agent.Coordinate{Latitude: 0, Longitude: 0.0001, Altitude: 30000}

	sim.Aircrafts[p1.Callsign] = p1
	sim.Aircrafts[p2.Callsign] = p2

	// Register agents to enable messaging and command application.
	sim.registerAircraftAgent(p1)
	sim.registerAircraftAgent(p2)

	// Run one simulation step to detect conflicts and emit directives.
	sim.step()

	ok := waitFor(t, 500*time.Millisecond, func() bool {
		// tower picks the first controlled aircraft (p1) and orders +1000 ft
		return p1.GetPosition().Altitude == 31000
	})
	if !ok {
		t.Fatalf("expected AFR100 altitude to change to 31000, got %d", p1.GetPosition().Altitude)
	}
}

func TestStopShutsDownCleanly(t *testing.T) {
	sim := NewSimulation()
	p := agent.NewAircraft("AFR300", "COMMERCIAL")
	sim.Aircrafts[p.Callsign] = p
	sim.registerAircraftAgent(p)

	done := make(chan struct{})
	go func() {
		sim.Stop()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("Stop did not return within timeout")
	}
}
