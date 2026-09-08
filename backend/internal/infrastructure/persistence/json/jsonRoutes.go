package json

// EnrouteRoute (Rte)

type Rte struct {
	RteUid string `json:"rteUid"`
	TxtRmk string `json:"txtRmk"`
}

type RteUid struct {
	Mid         int    `json:"mid"`
	TxtDesig    string `json:"txtDesig"`
	TxtLocDesig string `json:"txtLocDesig"`
}

// RouteSegment (Rsg)
type Rsg struct {
	CodeType           string `json:"codeType"`
	CodeRnp            int    `json:"codeRnp"`
	RsgUid             RsgUid `json:"rsgUid"`
	ValDistVerUpper    string `json:"valDistVerUpper"`
	ValDistVerLower    string `json:"valDistVerLower"`
	ValWid             int    `json:"valWid"`
	ValMagTrack        int    `json:"valMagTrack"`
	ValReverseMagTrack int    `json:"valReverseMagTrack"`
	ValLen             int    `json:"valLen"`
}

type RsgUid struct {
	Mid       int    `json:"mid"`
	RteUid    string `json:"rteUid"`
	DnpUidSta string `json:"dpnUidSta"`
	DpnUidEnd string `json:"dpnUidEnd"`
}

type DpnUidSta struct {
	Mid     int     `json:"mid"`
	CodeId  string  `json:"codeId"`
	GeoLat  float64 `json:"geoLat"`
	GeoLong float64 `json:"geoLong"`
}

type DpnUidEnd struct {
	Mid     int     `json:"mid"`
	CodeId  string  `json:"codeId"`
	GeoLat  float64 `json:"geoLat"`
	GeoLong float64 `json:"geoLong"`
}

// RouteSegmentUsage (Rsu)
type Rsu struct {
	RsuUid RsuUid `json:"rsuUid"`
	Rul    Rul    `json:"rul"`
}

type RsuUid struct {
	Mid         int    `json:"mid"`
	RsgUid      RsgUid `json:"rsgUid"`
	CodeRteAvbl string `json:"codeRteAvbl"`
	NoSeq       int    `json:"noSeq"`
	CodeDir     string `json:"codeDir"`
}

type Rul struct {
	PlcUid          string `json:"plcUid"`
	ValDistVerLower string `json:"valDistVerLower"`
	ValDistVerUpper string `json:"valDistVerUpper"`
}
