package json

// SignificantPointAirspace (Spa)
type Spa struct {
	SpaUid   SpaUid `json:"SpaUid"`
	CodeType string `json:"codeType"`
}

type SpaUid struct {
	Mid       int       `json:"mid"`
	DpnUidSpn DpnUidSpn `json:"dpnUidSpn"`
	AseUid    AseUid    `json:"aseUid"`
}

type DpnUidSpn struct {
	Mid     int     `json:"mid"`
	CodeId  string  `json:"codeId"`
	GeoLat  float64 `json:"geoLat"`
	GeoLong float64 `json:"geoLong"`
}

type Dpn struct {
	DpnUid      DpnUid      `json:"dpnUid"`
	AhpUidAssoc AhpUidAssoc `json:"ahpUidAssoc"`
	CodeDatum   string      `json:"codeDatum"`
	ValCrc      string      `json:"valCrc"`
	CodeType    string      `json:"codeType"`
	TxtName     string      `json:"txtName"`
}

type DpnUid struct {
	Mid     int     `json:"mid"`
	CodeId  string  `json:"codeId"`
	GeoLat  float64 `json:"geoLat"`
	GeoLong float64 `json:"geoLong"`
}

type AhpUidAssoc struct {
	Mid    int    `json:"mid"`
	CodeId string `json:"codeId"`
}
