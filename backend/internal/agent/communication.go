package agent

import (
	"context"
	"time"
)

// MessageType defines the nature of a communication exchanged between agents.
type MessageType string

const (
	MessageTypeConflictAlert       MessageType = "CONFLICT_ALERT"
	MessageTypeTowerDirective      MessageType = "TOWER_DIRECTIVE"
	MessageTypeAcknowledgement     MessageType = "ACKNOWLEDGEMENT"
	MessageTypeADSB                MessageType = "ADSB"
	MessageTypeArrivalNotification MessageType = "ARRIVAL_NOTIFICATION"
)

// AcknowledgementStatus provides more context for acknowledgement messages.
type AcknowledgementStatus string

const (
	AcknowledgementStatusAccepted AcknowledgementStatus = "ACCEPTED"
	AcknowledgementStatusRejected AcknowledgementStatus = "REJECTED"
)

// Message represents an envelope exchanged between simulation agents.
type Message struct {
	Type            MessageType
	From            string
	To              string
	Conflict        *ConflictAlert
	Directive       *TowerDirective
	Acknowledgement *Acknowledgement
	ADSB            *ADSBReport
	Arrival         *ArrivalNotification
}

// ConflictAlert carries information about a detected risk between two aircraft.
type ConflictAlert struct {
	ConflictID        string
	Aircraft1         string
	Aircraft2         string
	Level             string
	DistanceMeters    float64
	Aircraft1Altitude int
	Aircraft2Altitude int
}

// TowerDirective contains instructions emitted by a tower toward an aircraft.
type TowerDirective struct {
	ConflictID     string
	Instruction    string
	TargetAltitude int
}

// Acknowledgement reports the outcome of a directive execution.
type Acknowledgement struct {
	ConflictID string
	Status     AcknowledgementStatus
	Details    string
}

// CommandType lists the actions agents can request from the simulation engine.
type CommandType string

const (
	CommandTypeAdjustAltitude CommandType = "ADJUST_ALTITUDE"
)

// Command is the low-level instruction applied by the simulation engine.
type Command struct {
	Callsign       string
	Type           CommandType
	TargetAltitude int
	Notes          string
}

// ADSBReport transports positional data broadcast by an aircraft.
type ADSBReport struct {
	Emitter         string
	Position        Coordinate
	Speed           float64
	PrevWaypointID  string
	NextWaypointID  string
	DistanceToNext  float64
	SegmentProgress float64 // 0..1 along current segment
	Timestamp       time.Time
}

// ArrivalNotification signals that an aircraft has reached its destination.
type ArrivalNotification struct {
	AircraftCallsign string
	Destination      string
	ArrivalTime      time.Time
}

// Mailbox groups the channels used by an agent to communicate with the router.
type Mailbox struct {
	Inbox  <-chan Message
	Outbox chan<- Message
}

// Send forwards a message using the mailbox outbox channel with timeout to prevent deadlocks.
func (mb Mailbox) Send(msg Message) {
	if mb.Outbox == nil {
		return
	}
	select {
	case mb.Outbox <- msg:
		// Message envoyé avec succès
	case <-time.After(100 * time.Millisecond):
		// Drop message si timeout - acceptable pour éviter les deadlocks
	}
}

// SendWithContext forwards a message respecting context cancellation.
func (mb Mailbox) SendWithContext(ctx context.Context, msg Message) error {
	if mb.Outbox == nil {
		return nil
	}
	select {
	case mb.Outbox <- msg:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(100 * time.Millisecond):
		return ErrSendTimeout
	}
}

// ErrSendTimeout is returned when a message send operation times out.
var ErrSendTimeout = &TimeoutError{msg: "message send timeout"}

// TimeoutError represents a timeout during message sending.
type TimeoutError struct {
	msg string
}

func (e *TimeoutError) Error() string {
	return e.msg
}
