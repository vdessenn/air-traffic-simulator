package json

import "encoding/json"

// Organisation (Org)
type Org struct {
	OrgUid   OrgUid `json:"orgUid"`
	CodeId   string `json:"codeId"`
	CodeType string `json:"codeType"`
}

type OrgUid struct {
	Mid     int    `json:"mid"`
	TxtName string `json:"txtName"`
}

type SerUid struct {
	Mid      int             `json:"mid"`
	UniUid   json.RawMessage `json:"uniUid"`
	CodeType string          `json:"codeType"`
	NoSeq    int             `json:"noSeq"`
}

// UniteType (Uni)
type Uni struct {
	UniUid    UniUid  `json:"uniUid"`
	OrgUid    OrgUid  `json:"orgUid"`
	AhpUid    AhpUid  `json:"ahpUid"`
	CodeType  string  `json:"codeType"`
	CodeClass string  `json:"codeClass"`
	GeoLat    float64 `json:"geoLat"`
	GeoLong   float64 `json:"geoLong"`
	CodeDatum string  `json:"codeDatum"`
}

type UniUid struct {
	Mid     int    `json:"mid"`
	TxtName string `json:"txtName"`
}
