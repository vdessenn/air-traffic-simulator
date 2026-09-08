package aixm

import "encoding/xml"

// AirspaceType (Ase)
type Ase struct {
	XMLName          xml.Name `xml:"Ase"`
	AseUid           AseUid   `xml:"AseUid"`
	TxtName          string   `xml:"txtName"`
	CodeDistVerUpper string   `xml:"codeDistVerUpper"`
	ValDistVerUpper  string   `xml:"valDistVerUpper"`
	UomDistVerUpper  string   `xml:"uomDistVerUpper"`
	CodeDistVerLower string   `xml:"codeDistVerLower"`
	ValDistVerLower  string   `xml:"valDistVerLower"`
	UomDistVerLower  string   `xml:"uomDistVerLower"`
	Att              Att      `xml:"Att"`
	TxtRmk           string   `xml:"txtRmk"`
}

type AseUid struct {
	Mid      int    `xml:"mid,attr"`
	CodeType string `xml:"codeType"`
	CodeId   string `xml:"codeId"`
}

type Att struct {
	CodeWorkHr   string  `xml:"codeWorkHr"`
	Timsh        []Timsh `xml:"Timsh"` // plusieurs Timsh -> slice
	TxtRmkWorkHr string  `xml:"txtRmkWorkHr"`
}

type Timsh struct {
	CodeTimeRef  string `xml:"codeTimeRef"`
	DateValidWef string `xml:"dateValidWef"`
	DateValidTil string `xml:"dateValidTil"`
	CodeDay      string `xml:"codeDay"`
	CodeDayTil   string `xml:"codeDayTil"`
	TimeWef      string `xml:"timeWef"`
	TimeTil      string `xml:"timeTil"`
}

// AirspaceBorderType (Abd)
type Abd struct {
	XMLName xml.Name `xml:"Abd"`
	AbdUid  AbdUid   `xml:"AbdUid"`
	Avx     []Avx    `xml:"Avx"`
}

type AbdUid struct {
	Mid    int    `xml:"mid,attr"`
	AseUid AseUid `xml:"AseUid"`
}

type Avx struct {
	CodeType  string `xml:"codeType"`
	GeoLat    string `xml:"geoLat"`
	GeoLong   string `xml:"geoLong"`
	CodeDatum string `xml:"codeDatum"`
}

// AirspaceDerivedGeometryType (Adg)
type Adg struct {
	XMLName         xml.Name        `xml:"Adg"`
	AdgUid          AdgUid          `xml:"AdgUid"`
	AseUidBase      AseUidBase      `xml:"AseUidBase"`
	CodeOpr         string          `xml:"codeOpr"`
	AseUidComponent AseUidComponent `xml:"AseUidComponent"`
}

type AdgUid struct {
	AseUid AseUid `xml:"AseUid"`
}

type AseUidBase struct {
	Mid      int    `xml:"mid,attr"`
	CodeType string `xml:"codeType"`
	CodeId   string `xml:"codeId"`
}

type AseUidComponent struct {
	Mid      int    `xml:"mid,attr"`
	CodeType string `xml:"codeType"`
	CodeId   string `xml:"codeId"`
}

// FatoType (Fato) Final Approach and Take-Off Area
type Fto struct {
	XMLName         xml.Name `xml:"Fto"`
	FatoUid         FatoUid  `xml:"FatoUid"`
	ValLength       float64  `xml:"valLen"`
	ValWid          float64  `xml:"valWid"`
	UomDim          string   `xml:"uomDim"`
	CodeComposition string   `xml:"codeComposition"`
	ValSiwlWeight   float64  `xml:"valSiwlWeight"`
	UomSiwlWeight   string   `xml:"uomSiwlWeight"`
	TxtRmk          string   `xml:"txtRmk"`
}

type FatoUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

// GeographicalBorderType (Gbr)
type Gbr struct {
	XMLName  xml.Name `xml:"Gbr"`
	GbrUid   GbrUid   `xml:"GbrUid"`
	CodeType string   `xml:"codeType"`
	Gbv      []Gbv    `xml:"Gbv"`
}

type GbrUid struct {
	Mid     int    `xml:"mid,attr"`
	TxtName string `xml:"txtName"`
}

type Gbv struct {
	CodeType  string `xml:"codeType"`
	GeoLat    string `xml:"geoLat"`
	GeoLong   string `xml:"geoLong"`
	CodeDatum string `xml:"codeDatum"`
}

// AirspaceServiceType (Sae)
type Sae struct {
	XMLName xml.Name `xml:"Sae"`
	SaeUid  SaeUid   `xml:"SaeUid"`
}

type SaeUid struct {
	Mid    int    `xml:"mid,attr"`
	SerUid SerUid `xml:"SerUid"`
	AseUid AseUid `xml:"AseUid"`
}
