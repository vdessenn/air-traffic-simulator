package aixm

import "encoding/xml"

// ObstacleType (Obs)
type Obs struct {
	XMLName         xml.Name `xml:"Obs"`
	ObsUid          ObsUid   `xml:"ObsUid"`
	TxtName         string   `xml:"txtName"`
	TxtDescrType    string   `xml:"txtDescrType"`
	CodeGroup       string   `xml:"codeGroup"`
	CodeLgt         string   `xml:"codeLgt"`
	TxtDescrLgt     string   `xml:"txtDescrLgt"`
	ValGeoAccuracy  float64  `xml:"valGeoAccuracy"`
	UomGeoAccuracy  string   `xml:"uomGeoAccuracy"`
	ValElev         float64  `xml:"valElev"`
	ValElevAccuracy float64  `xml:"valElevAccuracy"`
	ValHgt          float64  `xml:"valHgt"`
	UomDistVer      string   `xml:"uomDistVer"`
	ValCrc          string   `xml:"valCrc"`
}

type ObsUid struct {
	Mid     int    `xml:"mid,attr"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}
