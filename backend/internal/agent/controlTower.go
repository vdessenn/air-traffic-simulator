package agent

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// TowerControl represente une tour de controle supervisant un ensemble d'aeronefs.
type TowerControl struct {
	BaseAgent

	mu                 sync.RWMutex
	ControlledAircraft []*Aircraft

	mailboxMu sync.RWMutex
	mailbox   Mailbox
}

// NewTowerControl construit un agent tour de controle avec un registre vide.
func NewTowerControl(callsign string) *TowerControl {
	return &TowerControl{
		BaseAgent:          NewBaseAgent(callsign),
		ControlledAircraft: []*Aircraft{},
	}
}

func (t *TowerControl) Act(ctx context.Context) error {
	for {
		t.mailboxMu.RLock()
		inbox := t.mailbox.Inbox
		t.mailboxMu.RUnlock()

		if inbox == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-inbox:
			if !ok {
				return nil
			}
			t.mailboxMu.RLock()
			mb := t.mailbox
			t.mailboxMu.RUnlock()
			t.handleMessage(ctx, msg, mb)
		default:
			return nil
		}
	}
}

// Run lance la boucle de traitement des messages pour la tour de controle.
func (t *TowerControl) Run(mailbox Mailbox) {
	t.mailboxMu.Lock()
	t.mailbox = mailbox
	t.mailboxMu.Unlock()
}

func (t *TowerControl) handleMessage(ctx context.Context, msg Message, mailbox Mailbox) {
	switch msg.Type {
	case MessageTypeADSB:
		// Tower doesn't process ADS-B messages, skip immediately
		return
	case MessageTypeConflictAlert:
		t.handleConflictAlert(ctx, msg, mailbox)
	case MessageTypeAcknowledgement:
		if msg.Acknowledgement != nil {
			log.Printf("[TOWER %s] Directive %s -> %s (%s)",
				t.Callsign,
				msg.Acknowledgement.ConflictID,
				msg.From,
				msg.Acknowledgement.Status,
			)
		}
	case MessageTypeArrivalNotification:
		t.handleArrivalNotification(ctx, msg)
	default:
		// message ignore
	}
}

func (t *TowerControl) handleConflictAlert(ctx context.Context, msg Message, mailbox Mailbox) {
	if msg.Conflict == nil {
		return
	}

	alert := msg.Conflict

	target, currentAltitude := t.chooseTargetAircraft(alert)
	if target == "" {
		return
	}

	targetAltitude := currentAltitude + 1000
	instruction := fmt.Sprintf("Adjust altitude to %d ft to resolve conflict %s with %s", targetAltitude, alert.ConflictID, otherAircraft(alert, target))

	directive := Message{
		Type: MessageTypeTowerDirective,
		From: t.Callsign,
		To:   target,
		Directive: &TowerDirective{
			ConflictID:     alert.ConflictID,
			Instruction:    instruction,
			TargetAltitude: targetAltitude,
		},
	}
	mailbox.Send(directive)
}

func otherAircraft(alert *ConflictAlert, target string) string {
	if target == alert.Aircraft1 {
		return alert.Aircraft2
	}
	return alert.Aircraft1
}

func (t *TowerControl) chooseTargetAircraft(alert *ConflictAlert) (string, int) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	for _, aircraft := range t.ControlledAircraft {
		callsign := aircraft.Callsign
		switch callsign {
		case alert.Aircraft1:
			return callsign, alert.Aircraft1Altitude
		case alert.Aircraft2:
			return callsign, alert.Aircraft2Altitude
		}
	}
	return "", 0
}

// ManageTakeoffLanding coordonne l'utilisation des pistes pour les departs et arrivees.
func (t *TowerControl) ManageTakeoffLanding() {
	// TODO: implementer la logique de planification des pistes.
}

// ManageAirspace garantit la separation des aeronefs et les limites de secteur.
func (t *TowerControl) ManageAirspace() {
	// TODO: implementer la logique de gestion de l'espace aerien.
}

// ManageCollisionRisks effectue la detection et la resolution des conflits.
func (t *TowerControl) ManageCollisionRisks() {
	// TODO: implementer la gestion des risques de collision.
}

// AssignAircraft ajoute un nouvel aeronef au registre de la tour s'il n'est pas deja present.
func (t *TowerControl) AssignAircraft(aircraft *Aircraft) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, controlled := range t.ControlledAircraft {
		if controlled == aircraft {
			return
		}
	}
	t.ControlledAircraft = append(t.ControlledAircraft, aircraft)
}

// ReleaseAircraft retire un aeronef du registre une fois transfere.
func (t *TowerControl) ReleaseAircraft(aircraft *Aircraft) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for idx, controlled := range t.ControlledAircraft {
		if controlled == aircraft {
			t.ControlledAircraft = append(t.ControlledAircraft[:idx], t.ControlledAircraft[idx+1:]...)
			return
		}
	}
}

// handleArrivalNotification processes arrival notifications and removes aircraft from controlled list.
func (t *TowerControl) handleArrivalNotification(ctx context.Context, msg Message) {
	if msg.Arrival == nil {
		return
	}

	log.Printf("[TOWER %s] Aircraft %s arrived at %s",
		t.Callsign,
		msg.Arrival.AircraftCallsign,
		msg.Arrival.Destination,
	)

	// Remove from controlled aircraft by callsign
	t.mu.Lock()
	defer t.mu.Unlock()

	for idx, aircraft := range t.ControlledAircraft {
		if aircraft.Callsign == msg.Arrival.AircraftCallsign {
			t.ControlledAircraft = append(t.ControlledAircraft[:idx], t.ControlledAircraft[idx+1:]...)
			log.Printf("[TOWER %s] Removed %s from controlled aircraft list",
				t.Callsign,
				msg.Arrival.AircraftCallsign,
			)
			return
		}
	}
}
