package aixm

import "encoding/xml"

// DesignatedPointType (Dpn)
type Dpn struct {
	XMLName     xml.Name    `xml:"Dpn"`
	DpnUid      DpnUid      `xml:"DpnUid"`
	AhpUidAssoc AhpUidAssoc `xml:"AhpUidAssoc"`
	CodeDatum   string      `xml:"codeDatum"`
	ValCrc      string      `xml:"valCrc"`
	CodeType    string      `xml:"codeType"`
	TxtName     string      `xml:"txtName"`
	TxtRmk      string      `xml:"txtRmk"`
}

type DpnUid struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type AhpUidAssoc struct {
	Mid    int    `xml:"mid,attr"`
	CodeId string `xml:"codeId"`
}

// SignificantPointAirspaceType (Spa)
type Spa struct {
	XMLName  xml.Name `xml:"Spa"`
	SpaUid   SpaUid   `xml:"SpaUid"`
	CodeType string   `xml:"codeType"`
	TxtRmk   string   `xml:"txtRmk"`
}

type SpaUid struct {
	Mid       int       `xml:"mid,attr"`
	DpnUidSpn DpnUidSpn `xml:"DpnUid"`
	AseUid    AseUid    `xml:"AseUid"`
}

type DpnUidSpn struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}
