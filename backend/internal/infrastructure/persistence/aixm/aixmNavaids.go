package aixm

import "encoding/xml"

// VorType (Vor)
type Vor struct {
	XMLName       xml.Name `xml:"Vor"`
	VorUid        VorUid   `xml:"VorUid"`
	OrgUid        OrgUid   `xml:"OrgUid"`
	TxtName       string   `xml:"txtName"`
	CodeType      string   `xml:"codeType"`
	ValFreq       float64  `xml:"valFreq"`
	UomFreq       string   `xml:"uomFreq"`
	CodeTypeNorth string   `xml:"codeTypeNorth"`
	ValElev       float64  `xml:"valElev"`
	UomDistVer    string   `xml:"uomDistVer"`
	ValCrc        string   `xml:"valCrc"`
	Vtt           Vtt      `xml:"Vtt"`
}

type VorUid struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type Vtt struct {
	CodeWorkHr   string `xml:"codeWord"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// NdbType (Ndb)
type Ndb struct {
	XMLName    xml.Name `xml:"Ndb"`
	NdbUid     NdbUid   `xml:"NdbUid"`
	OrgUid     OrgUid   `xml:"OrgUid"`
	TxtName    string   `xml:"txtName"`
	CodeType   string   `xml:"codeType"`
	ValFreq    float64  `xml:"valFreq"`
	UomFreq    string   `xml:"uomFreq"`
	CodeClass  string   `xml:"codeClass"`
	CodeEm     string   `xml:"codeEm"`
	CodeDatum  string   `xml:"codeDatum"`
	ValElev    float64  `xml:"valElev"`
	UomDistVer string   `xml:"uomDistVer"`
	ValCrc     string   `xml:"valCrc"`
	Ntt        Ntt      `xml:"Ntt"`
}

type NdbUid struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type Ntt struct {
	CodeWorkHr   string `xml:"codeWord"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// DmeType (Dme)
type Dme struct {
	XMLName      xml.Name `xml:"Dme"`
	DmeUid       DmeUid   `xml:"DmeUid"`
	OrgUid       OrgUid   `xml:"OrgUid"`
	VorUid       VorUid   `xml:"VorUid"`
	TxtName      string   `xml:"txtName"`
	CodeChannel  string   `xml:"codeChannel"`
	ValGhostFreq float64  `xml:"valGhostFreq"`
	UomGhostFreq string   `xml:"uomGhostFreq"`
	CodeDatum    string   `xml:"codeDatum"`
	ValElev      float64  `xml:"valElev"`
	UomDistVer   string   `xml:"uomDistVer"`
	ValCrc       string   `xml:"valCrc"`
	Dtt          Dtt      `xml:"Dtt"`
}

type DmeUid struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type Dtt struct {
	CodeWorkHr string `xml:"codeWord"`
}

// TacanType (Tcn)
type Tcn struct {
	XMLName     xml.Name `xml:"Tcn"`
	TcnUid      TcnUid   `xml:"TcnUid"`
	OrgUid      OrgUid   `xml:"OrgUid"`
	VorUid      VorUid   `xml:"VorUid"`
	TxtName     string   `xml:"txtName"`
	CodeChannel string   `xml:"codeChannel"`
	CodeDatum   string   `xml:"codeDatum"`
	Ttt         Ttt      `xml:"Ttt"`
}

type TcnUid struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type Ttt struct {
	CodeWorkHr string `xml:"codeWord"`
}

// IlsType (Ils)
type Ils struct {
	XMLName xml.Name `xml:"Ils"`
	IlsUid  IlsUid   `xml:"IlsUid"`
	DmeUid  DmeUid   `xml:"DmeUid"`
	CodeCat string   `xml:"codeCat"`
	Ilz     Ilz      `xml:"Ilz"`
	Igp     Igp      `xml:"Igp"`
}

type IlsUid struct {
	Mid    int    `xml:"mid,attr"`
	RdnUid RdnUid `xml:"RdnUid"`
}

type Ilz struct {
	CodeId     string  `xml:"codeId"`
	ValFreq    float64 `xml:"valFreq"`
	UomFreq    string  `xml:"uomFreq"`
	GeoLat     string  `xml:"geoLat"`
	GeoLong    string  `xml:"geoLong"`
	CodeDatum  string  `xml:"codeDatum"`
	ValElev    float64 `xml:"valElev"`
	UomDistVer string  `xml:"uomDistVer"`
	ValCrc     string  `xml:"valCrc"`
	Ilt        Ilt     `xml:"Ilt"`
}

type Ilt struct {
	CodeWorkHr string `xml:"codeWord"`
}

type Igp struct {
	ValFreq    float64 `xml:"valFreq"`
	UomFreq    string  `xml:"uomFreq"`
	ValSlope   float64 `xml:"valSlope"`
	ValRdh     float64 `xml:"valRdh"`
	GeoLat     string  `xml:"geoLat"`
	GeoLong    string  `xml:"geoLong"`
	CodeDatum  string  `xml:"codeDatum"`
	ValElev    float64 `xml:"valElev"`
	UomDistVer string  `xml:"uomDistVer"`
	ValCrc     string  `xml:"valCrc"`
	Igt        Igt     `xml:"Igt"`
}

type Igt struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// MkrType (Mkr)
type Mkr struct {
	XMLName    xml.Name `xml:"Mkr"`
	MkrUid     MkrUid   `xml:"MkrUid"`
	OrgUid     OrgUid   `xml:"OrgUid"`
	IlsUid     IlsUid   `xml:"IlsUid"`
	CodePsnIls string   `xml:"codePsnIls"`
	ValFreq    float64  `xml:"valFreq"`
	UomFreq    string   `xml:"uomFreq"`
	CodeDatum  string   `xml:"codeDatum"`
	ValElev    float64  `xml:"valElev"`
	UomDistVer string   `xml:"uomDistVer"`
	ValCrc     string   `xml:"valCrc"`
	Mtt        Mkt      `xml:"Mkt"`
}

type MkrUid struct {
	Mid     int    `xml:"mid,attr"`
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

type Mkt struct {
	CodeWorkHr string `xml:"codeWorkHr"`
}
