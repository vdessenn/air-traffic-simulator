package agent

import (
	"context"
	"sync"
)

// Agent definit le comportement que tous les agents autonomes doivent implementer.
type Agent interface {
	Percept(ctx context.Context) error
	Decide(ctx context.Context) error
	Act(ctx context.Context) error
	Communicate(ctx context.Context) error
	SetCommunicationUsed(commTypes []CommunicationType)
	CommunicationUsed() []CommunicationType
}

// BaseAgent regroupe les donnees et helpers communs aux implementations concretes.
type BaseAgent struct {
	Callsign           string
	communicationUsed  []CommunicationType
	communicationMutex sync.RWMutex
}

// NewBaseAgent instancie un BaseAgent avec le callsign fourni.
func NewBaseAgent(callsign string) BaseAgent {
	return BaseAgent{
		Callsign: callsign,
	}
}

// SetCommunicationUsed remplace la liste des types de communication autorises.
func (b *BaseAgent) SetCommunicationUsed(commTypes []CommunicationType) {
	b.communicationMutex.Lock()
	defer b.communicationMutex.Unlock()

	b.communicationUsed = make([]CommunicationType, len(commTypes))
	copy(b.communicationUsed, commTypes)
}

// CommunicationUsed renvoie les types de communication actuellement configures.
func (b *BaseAgent) CommunicationUsed() []CommunicationType {
	b.communicationMutex.RLock()
	defer b.communicationMutex.RUnlock()

	result := make([]CommunicationType, len(b.communicationUsed))
	copy(result, b.communicationUsed)

	return result
}
