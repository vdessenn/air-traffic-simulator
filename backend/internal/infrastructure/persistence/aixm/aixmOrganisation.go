package aixm

import "encoding/xml"

// OrganisationType (Org)
type Org struct {
	XMLName  xml.Name `xml:"Org"`
	OrgUid   OrgUid   `xml:"OrgUid"`
	CodeId   string   `xml:"codeId"`
	CodeType string   `xml:"codeType"`
}

type OrgUid struct {
	Mid     int    `xml:"mid,attr"`
	TxtName string `xml:"txtName"`
}

// FrequencyType (Fqy)
type Fqy struct {
	XMLName xml.Name `xml:"Fqy"`
	FqyUid  FqyUid   `xml:"fqyUid"`
	UomFreq string   `xml:"uomFreq"`
	Ftt     Ftt      `xml:"Ftt"`
	TxtRmk  string   `xml:"txtRmk"`
	Cdl     []Cdl    `xml:"Cdl"`
}

type FqyUid struct {
	Mid          int     `xml:"mid,attr"`
	SerUid       string  `xml:"serUid"`
	ValFreqTrans float64 `xml:"valFreqTrans"`
}

type Ftt struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

type Cdl struct {
	TxtCallSign string `xml:"txtCallSign"`
	CodeLang    string `xml:"codeLang"`
}

// ServiceType (Ser)
type Ser struct {
	XMLName   xml.Name `xml:"Ser"`
	SerUid    SerUid   `xml:"serUid"`
	GeoLat    string   `xml:"geoLat"`
	GeoLong   string   `xml:"geoLong"`
	CodeDatum string   `xml:"codeDatum"`
	ValCrc    string   `xml:"valCrc"`
	Stt       Stt      `xml:"Stt"`
	TxtRmk    string   `xml:"txtRmk"`
}

type SerUid struct {
	Mid      int    `xml:"mid,attr"`
	UniUid   UniUid `xml:"UniUid"`
	CodeType string `xml:"codeType"`
	NoSeq    int    `xml:"noSeq"`
}

type Stt struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// UnitType (Uni)
type Uni struct {
	XMLName   xml.Name `xml:"Uni"`
	UniUid    UniUid   `xml:"uniUid"`
	OrgUid    OrgUid   `xml:"OrgUid"`
	AhpUid    AhpUid   `xml:"AhpUid"`
	CodeType  string   `xml:"codeType"`
	CodeClass string   `xml:"codeClass"`
	GeoLat    string   `xml:"geoLat"`
	GeoLong   string   `xml:"geoLong"`
	CodeDatum string   `xml:"codeDatum"`
}

type UniUid struct {
	Mid     int    `xml:"mid,attr"`
	TxtName string `xml:"txtName"`
}
