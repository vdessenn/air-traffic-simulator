package json

// Airspace (Ase)
type Ase struct {
	AseUid           AseUid `json:"aseUid"`
	CodeDistVerUpper string `json:"codeDistVerUpper"`
	ValDistVerUpper  string `json:"valDistVerUpper"`
	UomDistVerUpper  string `json:"uomDistVerUpper"`
	CodeDistVerLower string `json:"codeDistVerLower"`
	ValDistVerLower  string `json:"valDistVerLower"`
	UomDistVerLower  string `json:"uomDistVerLower"`
}

type AseUid struct {
	Mid      int    `json:"mid"`
	CodeType string `json:"codeType"`
	CodeId   string `json:"codeId"`
}

// AirspaceBorder (Abd)
type Abd struct {
	AbdUid AbdUid `json:"abdUid"`
	Avx    []Avx  `json:"avx"`
}

type AbdUid struct {
	Mid    int    `json:"mid"`
	AseUid AseUid `json:"aseUid"`
}

type Avx struct {
	GeoLat  float64 `json:"geoLat"`
	GeoLong float64 `json:"geoLong"`
}

// AirspaceDerivedGeometry (Adg)
type Adg struct {
	AdgUid          AdgUid          `json:"adgUid"`
	AseUidBase      AseUidBase      `json:"aseUidBase"`
	CodeOpr         string          `json:"codeOpr"`
	AseUidComponent AseUidComponent `json:"aseUidComponent"`
}

type AdgUid struct {
	AseUid AseUid `json:"aseUid"`
}

type AseUidBase struct {
	Mid      int    `json:"mid"`
	CodeType string `json:"codeType"`
	CodeId   string `json:"codeId"`
}

type AseUidComponent struct {
	Mid      int    `json:"mid"`
	CodeType string `json:"codeType"`
	CodeId   string `json:"codeId"`
}

// GeographicalBorderType (Gbr)
type Gbr struct {
	GbrUid   GbrUid `json:"gbrUid"`
	CodeType string `json:"codeType"`
	Gbv      []Gbv  `json:"gbv"`
}

type GbrUid struct {
	Mid     int    `json:"mid"`
	TxtName string `json:"txtName"`
}

type Gbv struct {
	CodeType string  `json:"codeType"`
	GeoLat   float64 `json:"geoLat"`
	GeoLong  float64 `json:"geoLong"`
}

// AirspaceService (Sae)
type Sae struct {
	SaeUid SaeUid `json:"saeUid"`
}

type SaeUid struct {
	Mid    int    `json:"mid"`
	SerUid SerUid `json:"serUid"`
	AseUid AseUid `json:"aseUid"`
}
