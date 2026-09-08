package aixm

import "encoding/xml"

// RunwayType
type Rwy struct {
	XMLName                 xml.Name `xml:"Rwy"`
	RwyUid                  RwyUid   `xml:"rwyUid"`
	ValLen                  float64  `xml:"valLen"`
	UomDimRwy               string   `xml:"uomDimRwy"`
	CodeComposition         string   `xml:"codeComposition"`
	CodePcnPavementType     string   `xml:"codePcnPavementType"`
	CodePcnPavementSubgrade string   `xml:"codePcnPavementSubgrade"`
	CodePcnMaxTirePressure  string   `xml:"codePcnMaxTirePressure"`
	CodePcnEvalMethod       string   `xml:"codePcnEvalMethod"`
	TxtPcnNote              string   `xml:"txtPcnNote"`
	ValLenStrip             float64  `xml:"valLenStrip"`
	ValWidStrip             float64  `xml:"valWidStrip"`
	UomDimStrip             string   `xml:"uomDimStrip"`
}

type RwyUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

// RunwayCentreLinePositionType (Rcp)
type Rcp struct {
	XMLName    xml.Name `xml:"Rcp"`
	RcpUid     RcpUid   `xml:"RcpUid"`
	GeoLat     string   `xml:"geoLat"`
	GeoLong    string   `xml:"geoLong"`
	CodeDatum  string   `xml:"codeDatum"`
	ValTrueBrg float64  `xml:"valTrueBrg"`
	UomTrueBrg string   `xml:"uomTrueBrg"`
	ValMagBrg  float64  `xml:"valMagBrg"`
	UomMagBrg  string   `xml:"uomMagBrg"`
}

type RcpUid struct {
	Mid    int    `xml:"mid,attr"`
	RwyUid RwyUid `xml:"RwyUid"`
	NoEnd  int    `xml:"noEnd"`
}

// RunwayDirectionApproachLightingSystemType (Rda)
type Rda struct {
	XMLName  xml.Name `xml:"Rda"`
	RdaUid   RdaUid   `xml:"RdaUid"`
	CodeType string   `xml:"codeType"`
	Rdt      Rdt      `xml:"Rdt"`
}

type RdaUid struct {
	Mid    int    `xml:"mid,attr"`
	RwyUid RwyUid `xml:"RwyUid"`
	NoEnd  int    `xml:"noEnd"`
}

type Rdt struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// RunwayDirectionDeclaredDistanceType (Rdd)
type Rdd struct {
	XMLName xml.Name `xml:"Rdd"`
	RddUid  RddUid   `xml:"RddUid"`
	ValTora float64  `xml:"valTora"`
	UomTora string   `xml:"uomTora"`
	ValToda float64  `xml:"valToda"`
	UomToda string   `xml:"uomToda"`
	ValAsda float64  `xml:"valAsda"`
	UomAsda string   `xml:"uomAsda"`
	ValLda  float64  `xml:"valLda"`
	UomLda  string   `xml:"uomLda"`
}

type RddUid struct {
	Mid    int    `xml:"mid,attr"`
	RwyUid RwyUid `xml:"RwyUid"`
	NoEnd  int    `xml:"noEnd"`
}

// RunwayDirectionType (Rdn)
type Rdn struct {
	XMLName        xml.Name `xml:"Rdn"`
	RdnUid         RdnUid   `xml:"RdnUid"`
	CodeDesignator string   `xml:"codeDesignator"`
	ValMagHeading  string   `xml:"valMagHeading"`
	UomMagHeading  string   `xml:"uomMagHeading"`
	ValTrueHeading string   `xml:"valTrueHeading"`
	UomTrueHeading string   `xml:"uomTrueHeading"`
}

type RdnUid struct {
	Mid      int    `xml:"mid,attr"`
	RwyUid   RwyUid `xml:"RwyUid"`
	TxtDesig string `xml:"txtDesig"`
}

// RunwayDirectionLightingSystemType (Rls)
type Rls struct {
	XMLName  xml.Name `xml:"Rls"`
	RlsUid   RlsUid   `xml:"RlsUid"`
	CodeType string   `xml:"codeType"`
	Rlt      Rlt      `xml:"Rlt"`
}

type RlsUid struct {
	Mid    int    `xml:"mid,attr"`
	RwyUid RwyUid `xml:"RwyUid"`
	NoEnd  int    `xml:"noEnd"`
}

type Rlt struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
}

// RunwayProtectionAreaType (Rpa)
type Rpa struct {
	XMLName   xml.Name `xml:"Rpa"`
	RpaUid    RpaUid   `xml:"RpaUid"`
	ValWidth  float64  `xml:"valWidth"`
	UomWidth  string   `xml:"uomWidth"`
	ValLength float64  `xml:"valLength"`
	UomLength string   `xml:"uomLength"`
}

type RpaUid struct {
	Mid    int    `xml:"mid,attr"`
	RwyUid RwyUid `xml:"RwyUid"`
	NoEnd  int    `xml:"noEnd"`
}

// TlofType (Tla)
type Tla struct {
	XMLName            xml.Name `xml:"Tla"`
	TlaUid             TlaUid   `xml:"TlaUid"`
	FtoUid             FtoUid   `xml:"FtoUid"`
	GeoLat             string   `xml:"geoLat"`
	GeoLong            string   `xml:"geoLong"`
	CodeDatum          string   `xml:"codeDatum"`
	ValElev            float64  `xml:"valElev"`
	ValGeoidUndulation float64  `xml:"valGeoidUndulation"`
	UomDistVer         string   `xml:"uomDistVer"`
	ValCrc             string   `xml:"valCrc"`
	ValLen             float64  `xml:"valLen"`
	ValWid             float64  `xml:"valWid"`
	UomDim             string   `xml:"uomDim"`
	CodeComposition    string   `xml:"codeComposition"`
	ValSiwlWeight      float64  `xml:"valSiwlWeight"`
	UomSiwlWeight      string   `xml:"uomSiwlWeight"`
	TxtRmk             string   `xml:"txtRmk"`
}

type TlaUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

type FtoUid struct {
	Mid      int    `xml:"mid,attr"`
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

// CheckpointType (Nsc)
type Nsc struct {
	XMLName xml.Name `xml:"Nsc"`
	NscUid  NscUid   `xml:"NscUid"`
	TxtRmk  string   `xml:"txtRmk"`
}

type NscUid struct {
	Mid      int    `xml:"mid,attr"`
	GsdUid   GsdUid `xml:"GsdUid"`
	CodeType string `xml:"codeType"`
}
