package model

type RiskLevel string

const (
	RiskLevelCollision        RiskLevel = "COLLISION"
	RiskLevelRiskOfCollision  RiskLevel = "RISK"
	RiskLevelSafetyNotAssured RiskLevel = "WARNING"
)

type CollisionRisk struct {
	ID          string     `json:"id"`
	Aircraft1ID string     `json:"aircraft1Id"`
	Aircraft2ID string     `json:"aircraft2Id"`
	Level       RiskLevel  `json:"level"`
	Distance    float64    `json:"distance"`
	Message     string     `json:"message"`
	Location    Coordinate `json:"location"`
}
