package agent

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"
)

// Aircraft encapsule le comportement et l'etat d'un agent volant.
type Aircraft struct {
	BaseAgent
	Mu               sync.RWMutex // Protects all fields below - exported for engine package access
	Type             AircraftType
	Speed            float64
	Acceleration     float64
	ClimbRate        float64
	TurnRadius       float64
	Heading          int
	RemainingFuel    float64
	Autonomy         float64
	FlightTime       string
	BurnRate         float64
	Position         Coordinate
	NextWaypoint     Waypoint
	PreviousWaypoint Waypoint
	FlightPlan       []Waypoint

	// Segment tracking and ADS-B perception
	PrevWaypointID  string
	NextWaypointID  string
	DistanceToNext  float64
	SegmentProgress float64
	Neighbors       map[string]ADSBReport

	OptimalSpeed float64

	mailbox         Mailbox
	commandSink     chan<- Command
	SpeedRegEnabled bool
}

// NewAircraft construit une nouvelle instance d'aeronef avec ses caracteristiques de base.
func NewAircraft(callsign string, aircraftType AircraftType) *Aircraft {
	return &Aircraft{
		BaseAgent:    NewBaseAgent(callsign),
		Type:         aircraftType,
		Neighbors:    make(map[string]ADSBReport),
		OptimalSpeed: 460, // default cruise speed (kt)
	}
}

// Percept collecte les informations de l'environnement ; renvoie nil tant que non implemente.
func (a *Aircraft) Percept(ctx context.Context) error {
	// Perception is implicit: ADS-B messages populate the Neighbors map in Run().
	return nil
}

// Decide met a jour l'etat interne selon la derniere perception.
func (a *Aircraft) Decide(ctx context.Context) error {
	if a.SpeedRegEnabled {
		a.regulateSpeed()
	} else {
		a.handleMessages(ctx)
	}
	return nil
}

func (a *Aircraft) handleMessages(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-a.mailbox.Inbox:
			if !ok {
				return nil
			}
			a.handleMessage(ctx, msg)
		default:
			return nil
		}
	}
}

func (a *Aircraft) regulateSpeed() error {
	// Compute a feasible speed interval based on performance and simple direct-neighbor constraints.
	const sepNM = 5.0

	a.Mu.RLock()
	selfDist := a.DistanceToNext
	selfProgress := a.SegmentProgress
	prevSeg := a.PrevWaypointID
	nextSeg := a.NextWaypointID
	neighbors := make(map[string]ADSBReport, len(a.Neighbors))
	for k, v := range a.Neighbors {
		neighbors[k] = v
	}
	opt := a.OptimalSpeed
	a.Mu.RUnlock()

	if opt <= 0 {
		opt = 460
	}

	perfMin := opt * 0.94
	perfMax := opt * 1.03

	// Identify direct neighbors on same segment (prev->next).
	var pred *ADSBReport
	var succ *ADSBReport
	for _, n := range neighbors {
		if n.PrevWaypointID == prevSeg && n.NextWaypointID == nextSeg {
			if n.SegmentProgress > selfProgress {
				if pred == nil || n.SegmentProgress < pred.SegmentProgress {
					clone := n
					pred = &clone
				}
			} else if n.SegmentProgress < selfProgress {
				if succ == nil || n.SegmentProgress > succ.SegmentProgress {
					clone := n
					succ = &clone
				}
			}
		}
	}

	var vMin float64 = perfMin
	var vMax float64 = perfMax

	sepTime := sepNM / opt // hours

	if pred != nil && pred.Speed > 0 && selfDist > 0 {
		tPred := pred.DistanceToNext / pred.Speed
		maxSpeed := selfDist / (tPred + sepTime)
		if maxSpeed > 0 {
			vMax = math.Min(vMax, maxSpeed)
		}
	}

	if succ != nil && succ.Speed > 0 && selfDist > 0 {
		tSucc := succ.DistanceToNext / succ.Speed
		minSpeed := selfDist / (tSucc - sepTime)
		if minSpeed > 0 {
			vMin = math.Max(vMin, minSpeed)
		}
	}

	if vMin > vMax {
		// Interval infeasible: pick midpoint and clamp to perf bounds.
		chosen := (vMin + vMax) / 2
		if chosen < perfMin {
			chosen = perfMin
		}
		if chosen > perfMax {
			chosen = perfMax
		}
		a.SetSpeed(chosen)
		return nil
	}

	// Choose speed closest to optimal within [vMin, vMax].
	var chosen float64
	if opt < vMin {
		chosen = vMin
	} else if opt > vMax {
		chosen = vMax
	} else {
		chosen = opt
	}
	a.SetSpeed(chosen)
	return nil
}

// Act applique les decisions aux actionneurs de l'aeronef (bouchon pour l'instant).
func (a *Aircraft) Act(ctx context.Context) error {
	if len(a.FlightPlan) == 0 {
		return nil
	}

	targetWp := a.FlightPlan[0]
	var targetLat, targetLon float64

	targetLat = targetWp.Coordinate.Latitude
	targetLon = targetWp.Coordinate.Longitude

	dy := targetLat - a.Position.Latitude
	dx := targetLon - a.Position.Longitude
	distance := math.Sqrt(dx*dx + dy*dy)

	// Close enough
	if distance < 0.02 {
		a.FlightPlan = a.FlightPlan[1:]

		a.PreviousWaypoint = targetWp
		if len(a.FlightPlan) > 0 {
			a.NextWaypoint = a.FlightPlan[0]
		} else {
			a.Speed = 0
		}
		return nil
	}

	headingRad := math.Atan2(dx, dy)

	a.Heading = int(headingRad * 180 / math.Pi)
	if a.Heading < 0 {
		a.Heading += 360
	}

	return nil
}

// Communicate partage des mises a jour avec les autres agents.
func (a *Aircraft) Communicate(ctx context.Context) error {
	if a.mailbox.Outbox == nil {
		return nil
	}
	report := a.buildADSBReport()
	msg := Message{
		Type: MessageTypeADSB,
		From: a.Callsign,
		ADSB: &report,
	}
	a.mailbox.Send(msg)
	return nil
}

// Run demarre la boucle principale de communication de l'aeronef.
func (a *Aircraft) Start(mailbox Mailbox, commandSink chan<- Command) {
	a.mailbox = mailbox
	a.commandSink = commandSink
}

func (a *Aircraft) handleMessage(ctx context.Context, msg Message) {
	switch msg.Type {
	case MessageTypeTowerDirective:
		a.handleTowerDirective(ctx, msg, a.mailbox, a.commandSink)
	case MessageTypeConflictAlert:
		// Alertes informatives pour l'instant.
	case MessageTypeADSB:
		if msg.ADSB != nil && msg.ADSB.Emitter != a.Callsign {
			a.Mu.Lock()
			a.Neighbors[msg.ADSB.Emitter] = *msg.ADSB
			a.Mu.Unlock()
		}
	default:
		// Types de message non pris en charge.
	}
}

func (a *Aircraft) handleTowerDirective(ctx context.Context, msg Message, mailbox Mailbox, commandSink chan<- Command) {
	if msg.Directive == nil {
		return
	}

	targetAltitude := msg.Directive.TargetAltitude
	if targetAltitude <= 0 {
		targetAltitude = a.GetPosition().Altitude
	}

	if commandSink != nil {
		select {
		case <-ctx.Done():
			return
		case commandSink <- Command{
			Callsign:       a.Callsign,
			Type:           CommandTypeAdjustAltitude,
			TargetAltitude: targetAltitude,
			Notes:          msg.Directive.Instruction,
		}:
		case <-time.After(100 * time.Millisecond):
			// Command dropped due to timeout
			return
		}
	}

	ack := Message{
		Type: MessageTypeAcknowledgement,
		From: a.Callsign,
		To:   msg.From,
		Acknowledgement: &Acknowledgement{
			ConflictID: msg.Directive.ConflictID,
			Status:     AcknowledgementStatusAccepted,
			Details:    "Directive acknowledged",
		},
	}
	// Utiliser SendWithContext pour respecter l'annulation
	_ = mailbox.SendWithContext(ctx, ack)
}

// Navigate selectionne le prochain waypoint selon le plan de vol.
func (a *Aircraft) Navigate() {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	if len(a.FlightPlan) == 0 {
		return
	}

	a.PreviousWaypoint = a.NextWaypoint
	a.NextWaypoint = a.FlightPlan[0]
	a.FlightPlan = a.FlightPlan[1:]
}

// SetSpeed met a jour la vitesse actuelle de l'aeronef.
func (a *Aircraft) SetSpeed(newSpeed float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Speed = newSpeed
}

// SetAcceleration met a jour l'acceleration actuelle de l'aeronef.
func (a *Aircraft) SetAcceleration(newAcceleration float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Acceleration = newAcceleration
}

// SetClimbRate met a jour le taux de montee.
func (a *Aircraft) SetClimbRate(newClimbRate float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.ClimbRate = newClimbRate
}

// SetTurnRadius met a jour le rayon de virage.
func (a *Aircraft) SetTurnRadius(newTurnRadius float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.TurnRadius = newTurnRadius
}

// SetAltitude met a jour l'altitude dans la position de l'aeronef.
func (a *Aircraft) SetAltitude(newAltitude int) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Position.Altitude = newAltitude
}

// SetHeading met a jour le cap en degres.
func (a *Aircraft) SetHeading(newHeading int) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Heading = newHeading
}

// SetFlightTime met a jour la duree totale de vol.
func (a *Aircraft) SetFlightTime(newFlightTime string) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.FlightTime = newFlightTime
}

// SetBurnRate met a jour le debit de carburant.
func (a *Aircraft) SetBurnRate(newBurnRate float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.BurnRate = newBurnRate
}

// UpdateSegmentInfo safely updates the aircraft segment tracking.
func (a *Aircraft) UpdateSegmentInfo(prevID, nextID string, distanceToNext, progress float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.PrevWaypointID = prevID
	a.NextWaypointID = nextID
	a.DistanceToNext = distanceToNext
	a.SegmentProgress = progress
}

// GetSegmentInfo returns a snapshot of the segment tracking data.
func (a *Aircraft) GetSegmentInfo() (prevID, nextID string, distanceToNext, progress float64) {
	a.Mu.RLock()
	defer a.Mu.RUnlock()
	return a.PrevWaypointID, a.NextWaypointID, a.DistanceToNext, a.SegmentProgress
}

// AdjustTrajectory met a jour le prochain waypoint.
func (a *Aircraft) AdjustTrajectory(newWaypoint Waypoint) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.NextWaypoint = newWaypoint
}

// GetNextWaypoint renvoie la cible de navigation actuelle si elle existe.
func (a *Aircraft) GetNextWaypoint() (Waypoint, error) {
	a.Mu.RLock()
	defer a.Mu.RUnlock()

	if a.NextWaypoint.Name == "" && a.NextWaypoint.Coordinate == (Coordinate{}) {
		return Waypoint{}, errors.New("no next waypoint defined")
	}
	return a.NextWaypoint, nil
}

// SetRemainingFuel met a jour la quantite de carburant restante.
func (a *Aircraft) SetRemainingFuel(newRemainingFuel float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.RemainingFuel = newRemainingFuel
}

// SetAutonomy met a jour la duree maximale pendant laquelle l'aeronef peut voler.
func (a *Aircraft) SetAutonomy(newAutonomy float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Autonomy = newAutonomy
}

// MonitorFuel effectue une routine de surveillance du carburant en attente d'implementation.
func (a *Aircraft) MonitorFuel() {
	// TODO: surveiller l'utilisation du carburant avec BurnRate et FlightTime.
}

// EmergencyProcedures declenche une logique d'urgence temporaire.
func (a *Aircraft) EmergencyProcedures() {
	// TODO: implementer les procedures d'urgence (ex. rediriger vers la piste la plus proche).
}

// GetPosition returns a safe copy of the aircraft's current position.
func (a *Aircraft) GetPosition() Coordinate {
	a.Mu.RLock()
	defer a.Mu.RUnlock()
	return a.Position
}

// GetHeading returns the current heading of the aircraft.
func (a *Aircraft) GetHeading() int {
	a.Mu.RLock()
	defer a.Mu.RUnlock()
	return a.Heading
}

// GetSpeed returns the current speed of the aircraft.
func (a *Aircraft) GetSpeed() float64 {
	a.Mu.RLock()
	defer a.Mu.RUnlock()
	return a.Speed
}

// GetFlightPlanCopy returns a safe copy of the current flight plan.
func (a *Aircraft) GetFlightPlanCopy() []Waypoint {
	a.Mu.RLock()
	defer a.Mu.RUnlock()

	plan := make([]Waypoint, len(a.FlightPlan))
	copy(plan, a.FlightPlan)
	return plan
}

// buildADSBReport constructs an ADS-B snapshot for this aircraft.
func (a *Aircraft) buildADSBReport() ADSBReport {
	a.Mu.RLock()
	defer a.Mu.RUnlock()

	return ADSBReport{
		Emitter:         a.Callsign,
		Position:        a.Position,
		Speed:           a.Speed,
		PrevWaypointID:  a.PrevWaypointID,
		NextWaypointID:  a.NextWaypointID,
		DistanceToNext:  a.DistanceToNext,
		SegmentProgress: a.SegmentProgress,
		Timestamp:       time.Now(),
	}
}

// Shutdown cleanly terminates the aircraft agent and releases resources.
func (a *Aircraft) Shutdown() {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	// Clear neighbors map
	a.Neighbors = make(map[string]ADSBReport)

	// Clear flight plan
	a.FlightPlan = nil

	// Set speed to zero
	a.Speed = 0
}

// IsArrived checks if the aircraft has completed its flight plan and is ready for removal.
func (a *Aircraft) IsArrived() bool {
	a.Mu.RLock()
	defer a.Mu.RUnlock()

	const landingSpeedThreshold = 50.0 // knots
	hasPreviousWaypoint := a.PreviousWaypoint.Name != ""
	return len(a.FlightPlan) == 0 && a.Speed < landingSpeedThreshold && hasPreviousWaypoint
}
