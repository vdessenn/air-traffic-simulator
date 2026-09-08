package aixm

import (
	"encoding/xml"
)

// AerodromeHeliportType (Ahp)
type Ahp struct {
	XMLName            xml.Name `xml:"Ahp"`
	AhpUid             AhpUid   `xml:"AhpUid"`
	OrgUid             string   `xml:"OrgUid"`
	TxtName            string   `xml:"txtName"`
	CodeIcao           string   `xml:"codeIcao"`
	CodeIata           string   `xml:"codeIata"`
	CodeType           string   `xml:"codeType"`
	TxtDescrRefPt      string   `xml:"txtDescrRefPt"`
	GeoLat             string   `xml:"geoLat"`
	GeoLong            string   `xml:"geoLong"`
	CodeDatum          string   `xml:"codeDatum"`
	ValElev            float64  `xml:"valElev"`
	ValGeoidUndulation float64  `xml:"valGeoidUndulation"`
	UomDistVer         string   `xml:"uomDistVer"`
	TxtNameCitySer     string   `xml:"txtNameCitySer"`
	TxtDescrSite       string   `xml:"txtDescrSite"`
	ValMagVar          float64  `xml:"valMagVar"`
	DateMagVar         int      `xml:"dateMagVar"`
	ValMagVarChg       float64  `xml:"valMagVarChg"`
	ValRefT            float64  `xml:"valRefT"`
	UomRefT            string   `xml:"uomRefT"`
	TxtNameAdmin       string   `xml:"txtNameAdmin"`
	TxtDescrAcl        string   `xml:"txtDescrAcl"`
	ValTransitionAlt   float64  `xml:"valTransitionAlt"`
	UomTransitionAlt   string   `xml:"uomTransitionAlt"`
	Aht                struct {
		CodeWorkHr   string `xml:"codeWorkHr"`
		TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
	} `xml:"Aht"`
}

type AhpUid struct {
	Mid    int    `xml:"mid,attr"`
	CodeId string `xml:"codeId"`
}

// AerodromeHeliportAddressType (Aha)
type Aha struct {
	XMLName    xml.Name `xml:"Aha"`
	AhaUid     AhaUid   `xml:"ahaUid"`
	TxtAddress string   `xml:"txtAddress"`
}

type AhaUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	CodeType string `xml:"codeType"`
	NoSeq    int    `xml:"noSeq"`
}

// AerodromeHeliportObstacleType (Aho)
type Aho struct {
	XMLName xml.Name `xml:"Aho"`
	AhoUid  AhoUid   `xml:"ahoUid"`
}

type AhoUid struct {
	ObsUid ObsUid `xml:"ObsUid"`
	AhpUid AhpUid `xml:"AhpUid"`
}

// AeronauticalGroundLightType (Agl)
type Agl struct {
	XMLName   xml.Name `xml:"Agl"`
	AglUid    AglUid   `xml:"aglUid"`
	GeoLat    string   `xml:"geoLat"`
	GeoLong   string   `xml:"geoLong"`
	CodeDatum string   `xml:"codeDatum"`
	Agt       Agt      `xml:"Agt"`
}

type AglUid struct {
	Mid      int    `xml:"mid,attr"`
	TxtName  string `xml:"txtName"`
	CodeType string `xml:"codeType"`
}

type Agt struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// GroundServiceType (Ahs)
type Ahs struct {
	XMLName xml.Name `xml:"Ahs"`
	AhsUid  AhsUid   `xml:"ahsUid"`
	Ast     Ast      `xml:"Ast"`
}

type AhsUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	CodeType string `xml:"codeSeq"`
}

type Ast struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// AerodromeHeliportUsageType (Ahu)
type Ahu struct {
	XMLName         xml.Name          `xml:"Ahu"`
	AhuUid          AhuUid            `xml:"ahuUid"`
	UsageLimitation []UsageLimitation `xml:"usageLimitations"`
}

type AhuUid struct {
	Mid    int    `xml:"mid,attr"`
	AhpUid AhpUid `xml:"AhpUid"`
}

type UsageLimitation struct {
	CodeUsageLimitation string `xml:"codeUsageLimitation"`
	UsageCondition      struct {
		AircraftClass string `xml:"AircraftClass"`
		FlightClass   struct {
			CodeType    string `xml:"codeType"`
			CodeRule    string `xml:"codeRule"`
			CodeOrigin  string `xml:"codeOrigin"`
			CodePurpose string `xml:"codePurpose"`
		} `xml:"FlightClass"`
	} `xml:"usageCondition"`
}

// AerodromeHeliportServiceType (Sah)
type Sah struct {
	XMLName xml.Name `xml:"Sah"`
	SahUid  SahUid   `xml:"sahUid"`
}

type SahUid struct {
	Mid    int    `xml:"mid,attr"`
	AhpUid AhpUid `xml:"AhpUid"`
	SerUid SerUid `xml:"SerUid"`
}

// TaxiwayType (Twy)
type Twy struct {
	XMLName                 xml.Name `xml:"Twy"`
	TwyUid                  TwyUid   `xml:"twyUid"`
	CodeType                string   `xml:"codeType"`
	ValWid                  float64  `xml:"valWid"`
	UomDim                  string   `xml:"uomDim"`
	CodeComposition         string   `xml:"codeComposition"`
	ValPcnClass             float64  `xml:"valPcnClass"`
	CodePcnPavementType     string   `xml:"codePcnPavementType"`
	CodePcnPavementSubgrade string   `xml:"codePcnPavementSubgrade"`
	CodePcnMaxTirePressure  string   `xml:"codePcnMaxTirePressure"`
	CodeMPcnEvalMethod      string   `xml:"codeMPcnEvalMethod"`
}

type TwyUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

// StopwayType (Swy)
type Swy struct {
	XMLName xml.Name `xml:"Swy"`
	SwyUid  SwyUid   `xml:"swyUid"`
	ValLen  float64  `xml:"valLen"`
	UomDim  string   `xml:"uomDim"`
	CodeSts string   `xml:"codeSts"`
}

type SwyUid struct {
	Mid    int    `xml:"mid,attr"`
	RdnUid RdnUid `xml:"RdnUid"`
}

// ApronType (Apn)
type Apn struct {
	XMLName    xml.Name `xml:"Apn"`
	ApnUid     ApnUid   `xml:"ApnUid"`
	CodeSts    string   `xml:"codeSts"`
	TxtMarking string   `xml:"txtMarking"`
}

type ApnUid struct {
	Mid     int    `xml:"mid,attr"`
	AhpUid  AhpUid `xml:"AhpUid"`
	TxtName string `xml:"txtName"`
}

// GateStandType (Gsd)
type Gsd struct {
	XMLName          xml.Name `xml:"Gsd"`
	GsdUid           GsdUid   `xml:"GsdUid"`
	CodeType         string   `xml:"codeType"`
	TxtDescrRestrUse string   `xml:"txtDescrRestrUse"`
	GeoLat           string   `xml:"geoLat"`
	GeoLong          string   `xml:"geoLong"`
	CodeDatum        string   `xml:"codeDatum"`
	UomDistVer       string   `xml:"uomDistVer"`
	ValCrc           string   `xml:"valCrc"`
}

type GsdUid struct {
	Mid      int    `xml:"mid,attr"`
	ApnUid   ApnUid `xml:"ApnUid"`
	TxtDesig string `xml:"txtDesig"`
}

// PassengerFacilityType (Pfy)
type Pfy struct {
	XMLName  xml.Name `xml:"Pfy"`
	PfyUid   PfyUid   `xml:"PfyUid"`
	TxtDescr string   `xml:"txtDescr"`
	TxtRmk   string   `xml:"txtRmk"`
}

type PfyUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	CodeType string `xml:"codeType"`
	NoSeq    int    `xml:"noSeq"`
}
