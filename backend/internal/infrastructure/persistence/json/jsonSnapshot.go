package json

// JSON Snapshot Root Element

type AIMXSnapshot struct {
	Created string `json:"created"`
	Origin  string `json:"origin"`
	Version string `json:"version"`
	Ahp     []Ahp  `json:"Ahp"` // AerodromeHeliportType
	Ase     []Ase  `json:"Ase"` // AirspaceType
	Abd     []Abd  `json:"Abd"` // AirspaceBorderType
	Adg     []Adg  `json:"Adg"` // AirspaceDerivedGeometryType
	Gbr     []Gbr  `json:"Gbr"` // GeographicalBorderType
	Sae     []Sae  `json:"Sae"` // AirspaceServiceType
	Dpn     []Dpn  `json:"Dpn"` // DesignatedPointType
	Spa     []Spa  `json:"Spa"` // SignificantPointAirspaceType
	Rte     []Rte  `json:"Rte"` // RouteType
	Rsg     []Rsg  `json:"Rsg"` // RouteSegmentType
	Rsu     []Rsu  `json:"Rsu"` // RouteSegmentUsageType
	Org     []Org  `json:"Org"` // OrganisationAuthorityType
	Uni     []Uni  `json:"Uni"` // UnitType
	Avx     []Avx  `json:"Avx"` // AirspaceVertexType
	Gbv     []Gbv  `json:"Gbv"` // GeometricBufferVertexType
	Rul     []Rul  `json:"Rul"` // RouteUseDefinitionType
}
