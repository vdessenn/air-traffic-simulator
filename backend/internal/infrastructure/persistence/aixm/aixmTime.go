package aixm

import "encoding/xml"

// SpecialDateType (Spd)

type Spd struct {
	XMLName xml.Name `xml:"Spd"`
	SpdUid  SpdUid   `xml:"SpdUid"`
	TxtName string   `xml:"txtName"`
	TxtRmk  string   `xml:"txtRmk"`
}

type SpdUid struct {
	Mid      int    `xml:"mid,attr"`
	OrgUid   OrgUid `xml:"OrgUid"`
	CodeType string `xml:"codeType"`
	DateDay  string `xml:"dateDay"`
	DateYear string `xml:"dateYear"`
}
