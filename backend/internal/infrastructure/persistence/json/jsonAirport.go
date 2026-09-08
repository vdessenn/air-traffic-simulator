package json

// AerodromeHeliport (Ahp)

type Ahp struct {
	AhpUid           AhpUid  `json:"ahpUid"`
	OrgUid           string  `json:"orgUid"`
	GeoLat           float64 `json:"geoLat"`
	GeoLong          float64 `json:"geoLong"`
	CodeDatum        string  `json:"codeDatum"`
	ValTransitionAlt int     `json:"valTransitionAlt"`
	UomTransitionAlt string  `json:"uomTransitionAlt"`
}

type AhpUid struct {
	Mid    int    `json:"mid"`
	CodeId string `json:"codeId"`
}
