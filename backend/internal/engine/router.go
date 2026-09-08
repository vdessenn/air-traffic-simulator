package engine

import (
	"log"
	"time"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
)

func (s *Simulation) startRouter() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer s.shutdownMailboxes()

		for {
			select {
			case <-s.ctx.Done():
				log.Println("Router: draining remaining messages...")
				timeout := time.After(200 * time.Millisecond)
			drainLoop:
				for {
					select {
					case msg := <-s.messageBus:
						// Traiter les derniers messages ou les jeter selon le type
						if msg.Type != agent.MessageTypeADSB {
							s.routeMessage(msg)
						}
					case <-timeout:
						break drainLoop
					}
				}
				log.Println("Router: shutdown complete")
				return
			case msg := <-s.messageBus:
				s.routeMessage(msg)
			}
		}
	}()
}

func (s *Simulation) routeMessage(msg agent.Message) {
	switch msg.Type {
	case agent.MessageTypeADSB:
		s.routeADSB(msg)
	default:
		targets := make([]string, 0, 1)
		if msg.To != "" {
			targets = append(targets, msg.To)
		} else {
			s.commMu.RLock()
			for callsign := range s.towers {
				targets = append(targets, callsign)
			}
			s.commMu.RUnlock()
		}
		for _, to := range targets {
			s.commMu.RLock()
			inbox, ok := s.inboxes[to]
			s.commMu.RUnlock()
			if !ok {
				log.Printf("No inbox registered for %s, dropping message", to)
				continue
			}

			select {
			case <-s.ctx.Done():
				return
			case inbox <- msg:
			default:
				log.Printf("Inbox full for %s, dropping message (type=%s)", to, msg.Type)
			}
		}
	}
}

// routeADSB fan-outs ADS-B messages to all aircraft except the emitter.
func (s *Simulation) routeADSB(msg agent.Message) {
	if msg.ADSB == nil {
		return
	}

	// Get emitter position
	emitterPos := msg.ADSB.Position
	const rangeNM = 30.0 // Réduit à 30 NM pour limiter le nombre de messages avec beaucoup d'avions

	// Build list of receivers while holding the lock
	s.commMu.RLock()
	receivers := make(map[string]chan agent.Message, len(s.inboxes))
	for callsign, inbox := range s.inboxes {
		// Skip emitter
		if callsign == msg.From {
			continue
		}

		// Skip tower - it doesn't process ADS-B messages
		if _, isTower := s.towers[callsign]; isTower {
			continue
		}

		receivers[callsign] = inbox
	}
	s.commMu.RUnlock()

	// Check distance and send messages without holding commMu
	for callsign, inbox := range receivers {
		// Check if receiver is an aircraft and within range
		s.mu.RLock()
		receiverAircraft, exists := s.Aircrafts[callsign]
		s.mu.RUnlock()

		if !exists {
			// Aircraft was removed, skip
			continue
		}

		receiverPos := receiverAircraft.GetPosition()
		dist := CalculateDistance(
			emitterPos.Latitude, emitterPos.Longitude,
			receiverPos.Latitude, receiverPos.Longitude,
		)

		if dist > rangeNM {
			continue // Skip aircraft outside range
		}

		// Non-blocking send to aircraft inbox
		// Use select with default to avoid panic on closed channel
		select {
		case <-s.ctx.Done():
			return
		case inbox <- msg:
			// Sent successfully
		default:
			// Drop ADS-B message if receiver inbox is full or closed
		}
	}
}

// broadcastADSB publishes a snapshot for each aircraft.
func (s *Simulation) broadcastADSB() {
	// Build snapshot of aircraft references without holding lock during method calls
	s.mu.RLock()
	aircraftSnapshots := make([]*agent.Aircraft, 0, len(s.Aircrafts))
	for _, plane := range s.Aircrafts {
		aircraftSnapshots = append(aircraftSnapshots, plane)
	}
	s.mu.RUnlock() // Release lock BEFORE calling aircraft methods to prevent deadlock

	now := time.Now()
	reports := make([]agent.Message, 0, len(aircraftSnapshots))

	// Now iterate and call aircraft methods WITHOUT holding s.mu
	for _, plane := range aircraftSnapshots {
		// Aircraft may have been removed concurrently - this is safe
		prevID, nextID, distToNext, progress := plane.GetSegmentInfo()
		callsign := plane.Callsign
		report := agent.ADSBReport{
			Emitter:         callsign,
			Position:        plane.GetPosition(),
			Speed:           plane.GetSpeed(),
			PrevWaypointID:  prevID,
			NextWaypointID:  nextID,
			DistanceToNext:  distToNext,
			SegmentProgress: progress,
			Timestamp:       now,
		}
		msg := agent.Message{
			Type: agent.MessageTypeADSB,
			From: callsign,
			ADSB: &report,
		}
		reports = append(reports, msg)
	}

	// Send messages without holding any locks
	for _, msg := range reports {
		s.publishMessage(msg)
	}
}

func (s *Simulation) publishMessage(msg agent.Message) {
	// Distinction entre messages critiques et informatifs
	isCritical := msg.Type != agent.MessageTypeADSB

	if isCritical {
		// Messages critiques (conflits, directives) : envoi bloquant avec timeout court
		timer := time.NewTimer(100 * time.Millisecond)
		defer timer.Stop()

		select {
		case <-s.ctx.Done():
			return
		case s.messageBus <- msg:
			// Message envoyé avec succès
		case <-timer.C:
			// Timeout uniquement en cas de problème grave
			log.Printf("WARNING: Critical message dropped (type=%s)", msg.Type)
		}
	} else {
		// Messages informatifs (ADS-B) : drop immédiat si bus plein
		select {
		case <-s.ctx.Done():
			return
		case s.messageBus <- msg:
			// Message envoyé avec succès
		default:
			// Drop silencieux acceptable pour ADS-B
		}
	}
}

func (s *Simulation) startCommandWorker() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				// Drainer les commandes restantes avec timeout
				log.Println("CommandWorker: draining remaining commands...")
				timeout := time.After(100 * time.Millisecond)
			drainLoop:
				for {
					select {
					case cmd := <-s.commandQueue:
						s.applyCommand(cmd)
					case <-timeout:
						break drainLoop
					}
				}
				log.Println("CommandWorker: shutdown complete")
				return
			case cmd := <-s.commandQueue:
				s.applyCommand(cmd)
			}
		}
	}()
}

func (s *Simulation) shutdownMailboxes() {
	s.commMu.Lock()
	defer s.commMu.Unlock()

	for _, inbox := range s.inboxes {
		close(inbox)
	}

	s.inboxes = make(map[string]chan agent.Message)
}

func (s *Simulation) applyCommand(cmd agent.Command) {
	s.mu.Lock()
	defer s.mu.Unlock()

	plane, ok := s.Aircrafts[cmd.Callsign]
	if !ok {
		return
	}

	switch cmd.Type {
	case agent.CommandTypeAdjustAltitude:
		plane.SetAltitude(cmd.TargetAltitude)
	}
}
