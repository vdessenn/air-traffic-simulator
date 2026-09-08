package aixm

import "encoding/xml"

// EnrouteRouteType (Rte)
type Rte struct {
	XMLName xml.Name `xml:"Rte"`
	RteUid  RteUid   `xml:"RteUid"`
	TxtRmk  string   `xml:"txtRmk"`
}

type RteUid struct {
	Mid         int    `xml:"mid,attr"`
	TxtDesig    string `xml:"txtDesig"`
	TxtLocDesig string `xml:"txtLocDesig"`
}

// RouteSegmentType (Rsg)
type Rsg struct {
	XMLName            xml.Name `xml:"Rsg"`
	RsgUid             RsgUid   `xml:"RsgUid"`
	CodeType           string   `xml:"codeType"`
	CodeRnp            float64  `xml:"codeRnp"`
	CodeLvl            string   `xml:"codeLvl"`
	ValDistVerUpper    string   `xml:"valDistVerUpper"`
	UomDistVerUpper    string   `xml:"uomDistVerUpper"`
	CodeDistVerUpper   string   `xml:"codeDistVerUpper"`
	ValDistVerLower    string   `xml:"valDistVerLower"`
	UomDistVerLower    string   `xml:"uomDistVerLower"`
	CodeDistVerLower   string   `xml:"codeDistVerLower"`
	ValWid             float64  `xml:"valWid"`
	UomWid             string   `xml:"uomWid"`
	CodeRepAtcStart    string   `xml:"codeRepAtcStart"`
	CodeRepAtcEnd      string   `xml:"codeRepAtcEnd"`
	CodeTypePath       string   `xml:"codeTypePath"`
	ValMagTrack        float64  `xml:"valMagTrack"`
	ValReverseMagTrack float64  `xml:"valReverseMagTrack"`
	ValLen             float64  `xml:"valLen"`
	UomDist            string   `xml:"uomDist"`
	TxtRmk             string   `xml:"txtRmk"`
}

type RsgUid struct {
	Mid       int       `xml:"mid,attr"`
	RteUid    RteUid    `xml:"RteUid"`
	DnpUidSta DpnUidSta `xml:"DpnUidSta"`
	DpnUidEnd DpnUidEnd `xml:"DpnUidEnd"`
}

type DpnUidSta struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type DpnUidEnd struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

// RouteSegmentUsageType (Rsu)
type Rsu struct {
	XMLName xml.Name `xml:"Rsu"`
	RsuUid  RsuUid   `xml:"RsuUid"`
	Rul     Rul      `xml:"Rul"`
	Rst     Rst      `xml:"Rst"`
}

type RsuUid struct {
	Mid         int    `xml:"mid,attr"`
	Rsg         RsgUid `xml:"Rsg"`
	CodeRteAvbl string `xml:"codeRteAvbl"`
	NoSeq       int    `xml:"noSeq"`
	CodeDir     string `xml:"codeDir"`
}

type Rul struct {
	PlcUid           PlcUid `xml:"PlcUid"`
	ValDistVerLower  string `xml:"valDistVerLower"`
	UomDistVerLower  string `xml:"uomDistVerLower"`
	CodeDistVerLower string `xml:"codeDistVerLower"`
	ValDistVerUpper  string `xml:"valDistVerUpper"`
	UomDistVerUpper  string `xml:"uomDistVerUpper"`
	CodeDistVerUpper string `xml:"codeDistVerUpper"`
}

type Rst struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// CruisingLevelsTableType (Plb)
type Plb struct {
	XMLName     xml.Name `xml:"Plb"`
	PlbUid      PlbUid   `xml:"PlbUid"`
	CodeDistVer string   `xml:"codeDistVer"`
	UomDistVer  string   `xml:"uomDistVer"`
}

type PlbUid struct {
	Mid    int    `xml:"mid,attr"`
	CodeId string `xml:"codeId"`
}

// CruisingLevelsColumnType (Plc)
type Plc struct {
	XMLName xml.Name `xml:"Plc"`
	PlcUid  PlcUid   `xml:"PlcUid"`
	Pll     Pll      `xml:"Pll"`
}

type PlcUid struct {
	Mid    int    `xml:"mid,attr"`
	PlbUid PlbUid `xml:"PlbUid"`
	CodeId string `xml:"codeId"`
}

type Pll struct {
	ValdDistVer string `xml:"valDistVer"`
}
