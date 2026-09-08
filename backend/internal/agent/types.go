package agent

// CommunicationType definit un support de communication disponible pour les agents.
type CommunicationType struct {
	Name      string
	Frequency float64
	Range     float64
}

// Coordinate modele la position 3D d'un objet dans l'espace de simulation.
type Coordinate struct {
	Latitude  float64
	Longitude float64
	Altitude  int
}

// Waypoint represente une cible de navigation dans le plan de vol.
type Waypoint struct {
	Name       string
	Coordinate Coordinate
}

// AircraftType identifie la categorie d'un aeronef.
type AircraftType string
