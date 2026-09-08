package engine

import (
	"testing"
	"time"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
)

func TestSimulationAircraftRemovalOnArrival(t *testing.T) {
	// Create simulation
	sim := NewSimulation()
	defer sim.Stop()

	// Create an aircraft that will arrive soon
	aircraft := agent.NewAircraft("TEST_ARRIVAL", agent.AircraftType("B737"))
	aircraft.Position = agent.Coordinate{
		Latitude:  48.0,
		Longitude: 2.0,
		Altitude:  1000,
	}
	aircraft.FlightPlan = []agent.Waypoint{} // Empty flight plan
	aircraft.SetSpeed(30.0)                  // Low speed indicating landing
	aircraft.PreviousWaypoint = agent.Waypoint{
		Name: "LFPG",
		Coordinate: agent.Coordinate{
			Latitude:  48.0,
			Longitude: 2.0,
		},
	}

	// Register aircraft in simulation
	sim.mu.Lock()
	sim.Aircrafts["TEST_ARRIVAL"] = aircraft
	sim.mu.Unlock()

	// Start the aircraft agent
	inbox := make(chan agent.Message, 100)
	outbox := make(chan agent.Message, 100)
	mailbox := agent.Mailbox{Inbox: inbox, Outbox: outbox}

	sim.commMu.Lock()
	sim.inboxes["TEST_ARRIVAL"] = inbox
	sim.commMu.Unlock()

	aircraft.Start(mailbox, sim.commandQueue)

	// Get initial aircraft count
	initialCount := len(sim.GetAircrafts())
	if initialCount != 1 {
		t.Fatalf("Expected 1 aircraft initially, got %d", initialCount)
	}

	// Run simulation step - should detect arrival and remove aircraft
	sim.step()

	// Give some time for async operations
	time.Sleep(100 * time.Millisecond)

	// Verify aircraft was removed
	finalCount := len(sim.GetAircrafts())
	if finalCount != 0 {
		t.Errorf("Expected 0 aircraft after arrival, got %d", finalCount)
	}

	// Verify aircraft no longer in registry
	sim.mu.RLock()
	_, exists := sim.Aircrafts["TEST_ARRIVAL"]
	sim.mu.RUnlock()

	if exists {
		t.Errorf("Aircraft should have been removed from registry")
	}

	// Verify inbox was closed and removed
	sim.commMu.RLock()
	_, inboxExists := sim.inboxes["TEST_ARRIVAL"]
	sim.commMu.RUnlock()

	if inboxExists {
		t.Errorf("Aircraft inbox should have been removed")
	}
}

func TestSimulationMultipleArrivals(t *testing.T) {
	sim := NewSimulation()
	defer sim.Stop()

	// Create multiple aircraft - some arriving, some not
	arrivals := []string{"ARRIVAL1", "ARRIVAL2"}
	flying := []string{"FLYING1", "FLYING2"}

	// Setup arriving aircraft
	for _, callsign := range arrivals {
		aircraft := agent.NewAircraft(callsign, agent.AircraftType("B737"))
		aircraft.Position = agent.Coordinate{Latitude: 48.0, Longitude: 2.0, Altitude: 500}
		aircraft.FlightPlan = []agent.Waypoint{}
		aircraft.SetSpeed(25.0)
		aircraft.PreviousWaypoint = agent.Waypoint{Name: "DEST"}

		sim.mu.Lock()
		sim.Aircrafts[callsign] = aircraft
		sim.mu.Unlock()

		inbox := make(chan agent.Message, 100)
		outbox := make(chan agent.Message, 100)
		mailbox := agent.Mailbox{Inbox: inbox, Outbox: outbox}

		sim.commMu.Lock()
		sim.inboxes[callsign] = inbox
		sim.commMu.Unlock()

		aircraft.Start(mailbox, sim.commandQueue)
	}

	// Setup flying aircraft
	for _, callsign := range flying {
		aircraft := agent.NewAircraft(callsign, agent.AircraftType("B737"))
		aircraft.Position = agent.Coordinate{Latitude: 49.0, Longitude: 3.0, Altitude: 35000}
		aircraft.FlightPlan = []agent.Waypoint{
			{Name: "WP1", Coordinate: agent.Coordinate{Latitude: 50.0, Longitude: 4.0}},
		}
		aircraft.SetSpeed(450.0)

		sim.mu.Lock()
		sim.Aircrafts[callsign] = aircraft
		sim.mu.Unlock()

		inbox := make(chan agent.Message, 100)
		outbox := make(chan agent.Message, 100)
		mailbox := agent.Mailbox{Inbox: inbox, Outbox: outbox}

		sim.commMu.Lock()
		sim.inboxes[callsign] = inbox
		sim.commMu.Unlock()

		aircraft.Start(mailbox, sim.commandQueue)
	}

	// Initial count should be 4
	initialCount := len(sim.GetAircrafts())
	if initialCount != 4 {
		t.Fatalf("Expected 4 aircraft initially, got %d", initialCount)
	}

	// Run simulation step
	sim.step()

	// Give time for async operations
	time.Sleep(100 * time.Millisecond)

	// Should have 2 aircraft remaining (the flying ones)
	finalCount := len(sim.GetAircrafts())
	if finalCount != 2 {
		t.Errorf("Expected 2 aircraft remaining, got %d", finalCount)
	}

	// Verify specific aircraft
	sim.mu.RLock()
	for _, callsign := range arrivals {
		if _, exists := sim.Aircrafts[callsign]; exists {
			t.Errorf("Aircraft %s should have been removed", callsign)
		}
	}
	for _, callsign := range flying {
		if _, exists := sim.Aircrafts[callsign]; !exists {
			t.Errorf("Aircraft %s should still be flying", callsign)
		}
	}
	sim.mu.RUnlock()
}

func TestTowerArrivalNotification(t *testing.T) {
	sim := NewSimulation()
	defer sim.Stop()

	// Get the default tower
	sim.mu.RLock()
	tower, exists := sim.towers[defaultTowerCallsign]
	sim.mu.RUnlock()

	if !exists {
		t.Fatal("Default tower should exist")
	}

	// Create and register an aircraft
	aircraft := agent.NewAircraft("TEST_NOTIFY", agent.AircraftType("B737"))
	aircraft.Position = agent.Coordinate{Latitude: 48.0, Longitude: 2.0, Altitude: 100}
	aircraft.FlightPlan = []agent.Waypoint{}
	aircraft.SetSpeed(20.0)
	aircraft.PreviousWaypoint = agent.Waypoint{Name: "LFPG"}

	sim.mu.Lock()
	sim.Aircrafts["TEST_NOTIFY"] = aircraft
	sim.mu.Unlock()

	inbox := make(chan agent.Message, 100)
	outbox := make(chan agent.Message, 100)
	mailbox := agent.Mailbox{Inbox: inbox, Outbox: outbox}

	sim.commMu.Lock()
	sim.inboxes["TEST_NOTIFY"] = inbox
	sim.commMu.Unlock()

	aircraft.Start(mailbox, sim.commandQueue)

	// Assign to tower
	tower.AssignAircraft(aircraft)

	// Run simulation step - should trigger arrival and notification
	sim.step()

	// Give time for message processing
	time.Sleep(200 * time.Millisecond)

	// Verify aircraft was removed from simulation
	finalCount := len(sim.GetAircrafts())
	if finalCount != 0 {
		t.Errorf("Expected 0 aircraft in simulation after arrival, got %d", finalCount)
	}
}
