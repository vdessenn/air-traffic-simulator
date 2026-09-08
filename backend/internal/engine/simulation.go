package engine

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/model"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/model/object"
)

const defaultTowerCallsign = "GLOBAL_TWR"

type Simulation struct {
	mu            sync.RWMutex
	Aircrafts     map[string]*agent.Aircraft
	Waypoints     []object.Waypoint
	Risks         []model.CollisionRisk
	WaypointsMap  map[string]model.Coordinate
	Sectors       []*model.Sector
	RouteSegments []RouteSegment
	//
	SpeedFactor    float64
	EnableSpeedReg bool
	ctx            context.Context
	cancel         context.CancelFunc
	//
	commandQueue chan agent.Command
	messageBus   chan agent.Message
	//
	commMu  sync.RWMutex
	inboxes map[string]chan agent.Message
	towers  map[string]*agent.TowerControl
	wg      sync.WaitGroup
	//
	activeConflicts map[string]time.Time
	tickCounter     int // Compteur de ticks pour limiter broadcasts
}

// RouteSegment models a simple directed edge between two waypoints.
type RouteSegment struct {
	FromID string
	ToID   string
	Length float64 // nautical miles
}

func NewSimulation() *Simulation {
	ctx, cancel := context.WithCancel(context.Background())

	sim := &Simulation{
		Aircrafts:       make(map[string]*agent.Aircraft),
		Waypoints:       make([]object.Waypoint, 0),
		Risks:           make([]model.CollisionRisk, 0),
		WaypointsMap:    make(map[string]model.Coordinate),
		Sectors:         make([]*model.Sector, 0),
		SpeedFactor:     1.0,
		EnableSpeedReg:  true, // enable speed-based regulation by default
		ctx:             ctx,
		cancel:          cancel,
		commandQueue:    make(chan agent.Command, 4096),
		messageBus:      make(chan agent.Message, 65536), // Très large buffer pour 120+ avions (chaque avion broadcast à ~120 voisins)
		inboxes:         make(map[string]chan agent.Message),
		towers:          make(map[string]*agent.TowerControl),
		activeConflicts: make(map[string]time.Time),
	}

	sim.startCommandWorker()
	sim.startRouter()

	defaultTower := agent.NewTowerControl(defaultTowerCallsign)
	sim.registerTowerAgent(defaultTower)

	return sim
}

func (s *Simulation) Start() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.step()
		}
	}
}

func (s *Simulation) Stop() {
	s.cancel()

	time.Sleep(50 * time.Millisecond)

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("Simulation stopped cleanly")
	case <-time.After(5 * time.Second):
		log.Println("WARNING: Graceful shutdown timeout - some goroutines may not have terminated")
	}
}

func (s *Simulation) step() {
	var alerts []agent.Message
	var arrivedAircraftToCleanup []*agent.Aircraft

	s.mu.Lock()

	s.updateAircraftPhysics()
	s.updateSectorOccupancy()
	s.Risks = DetectConflicts(s.Aircrafts)

	// Detect and handle aircraft arrivals
	// Remove from registry while holding lock, cleanup later without lock
	arrivedAircraft := DetectArrivals(s.Aircrafts)
	for _, callsign := range arrivedAircraft {
		if aircraft := s.removeAircraftFromRegistryLocked(callsign); aircraft != nil {
			arrivedAircraftToCleanup = append(arrivedAircraftToCleanup, aircraft)
		}
	}

	current := make(map[string]struct{}, len(s.Risks))
	for _, risk := range s.Risks {
		current[risk.ID] = struct{}{}
		if _, known := s.activeConflicts[risk.ID]; !known {
			alerts = append(alerts, s.buildConflictAlertLocked(risk))
			s.activeConflicts[risk.ID] = time.Now()
		}
	}

	for id := range s.activeConflicts {
		if _, ok := current[id]; !ok {
			delete(s.activeConflicts, id)
		}
	}

	for _, aircraft := range s.Aircrafts {
		aircraft.Percept(s.ctx)
		aircraft.Decide(s.ctx)
		aircraft.Act(s.ctx)
	}

	for _, tower := range s.towers {
		tower.Act(s.ctx)
	}

	if len(s.Risks) > 0 {
		log.Printf("Alert / %d conflict(s): ", len(s.Risks))
	}

	s.mu.Unlock()

	// Cleanup arrived aircraft WITHOUT holding simulation lock
	for _, aircraft := range arrivedAircraftToCleanup {
		s.cleanupAircraft(aircraft)
	}

	for _, msg := range alerts {
		s.publishMessage(msg)
	}

	// Limiter les broadcasts ADS-B selon le SpeedFactor pour éviter surcharge
	// À vitesse normale (1x): broadcast tous les 5 ticks = 500ms
	// À vitesse x10: broadcast tous les 5 ticks = toujours 500ms temps réel
	s.tickCounter++
	broadcastInterval := 5 // Base: tous les 5 ticks
	if s.SpeedFactor >= 5.0 {
		// À haute vitesse, espacer encore plus pour réduire la charge
		broadcastInterval = int(10 * s.SpeedFactor / 5.0)
	}
	if s.tickCounter%broadcastInterval == 0 {
		s.broadcastADSB()
	}
}

func (s *Simulation) GetAircrafts() []*agent.Aircraft {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]*agent.Aircraft, 0, len(s.Aircrafts))
	for _, aircraft := range s.Aircrafts {
		list = append(list, aircraft)
	}
	return list
}

func (s *Simulation) GetWaypoints() []object.Waypoint {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Waypoints
}

func (s *Simulation) GetRisks() []model.CollisionRisk {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]model.CollisionRisk, len(s.Risks))
	copy(result, s.Risks)
	return result
}

func (s *Simulation) GetRisksSummary() [3]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := [3]int{0, 0, 0}
	for _, risk := range s.Risks {
		switch risk.Level {
		case model.RiskLevelSafetyNotAssured:
			result[0] += 1
		case model.RiskLevelRiskOfCollision:
			result[1] += 1
		case model.RiskLevelCollision:
			result[2] += 1
		default:
		}
	}
	return result
}

func (s *Simulation) GetSectors() []*model.Sector {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Sectors
}

// removeAircraftFromRegistryLocked removes an aircraft from the registry and returns it for cleanup.
// MUST be called while holding s.mu.Lock()
// Returns the aircraft pointer if it existed, nil otherwise.
func (s *Simulation) removeAircraftFromRegistryLocked(callsign string) *agent.Aircraft {
	aircraft, exists := s.Aircrafts[callsign]
	if !exists {
		return nil
	}

	// Get destination for notification (last waypoint if available)
	destination := "UNKNOWN"
	if aircraft.PreviousWaypoint.Name != "" {
		destination = aircraft.PreviousWaypoint.Name
	}

	log.Printf("Aircraft %s arrived at %s - removing from simulation", callsign, destination)

	// Remove from registry to prevent new references
	delete(s.Aircrafts, callsign)

	return aircraft
}

// cleanupAircraft performs cleanup operations on a removed aircraft.
// MUST be called WITHOUT holding s.mu - handles its own locking as needed.
func (s *Simulation) cleanupAircraft(aircraft *agent.Aircraft) {
	if aircraft == nil {
		return
	}

	callsign := aircraft.Callsign
	destination := "UNKNOWN"
	if aircraft.PreviousWaypoint.Name != "" {
		destination = aircraft.PreviousWaypoint.Name
	}

	// Cleanup aircraft resources without holding simulation lock
	// This prevents deadlock with concurrent operations that hold aircraft.Mu
	aircraft.Shutdown()

	// Remove inbox - must be done carefully to avoid race with router
	// We don't close the channel immediately to avoid panic on send to closed channel
	// Instead, we remove it from the map and let it be garbage collected
	s.commMu.Lock()
	if inbox, ok := s.inboxes[callsign]; ok {
		delete(s.inboxes, callsign)
		// Drain the inbox in a goroutine to prevent blocking and close it safely
		go func(ch chan agent.Message) {
			// Give time for any in-flight messages
			time.Sleep(50 * time.Millisecond)
			// Drain remaining messages
			for len(ch) > 0 {
				<-ch
			}
			close(ch)
		}(inbox)
	}
	s.commMu.Unlock()

	// Notify control tower with timeout to prevent blocking
	arrivalNotification := agent.Message{
		Type: agent.MessageTypeArrivalNotification,
		From: "SIMULATION",
		To:   defaultTowerCallsign,
		Arrival: &agent.ArrivalNotification{
			AircraftCallsign: callsign,
			Destination:      destination,
			ArrivalTime:      time.Now(),
		},
	}

	// Send notification with timeout to avoid blocking cleanup
	select {
	case s.messageBus <- arrivalNotification:
		// Sent successfully
	case <-time.After(100 * time.Millisecond):
		// Timeout - acceptable during shutdown or high load
	case <-s.ctx.Done():
		// Simulation stopping
	}
}
