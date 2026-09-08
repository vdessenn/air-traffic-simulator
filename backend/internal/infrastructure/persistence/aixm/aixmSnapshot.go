package aixm

import "encoding/xml"

// AIXM Snapshot Root Element
type AIXMSnapshot struct {
	XMLName xml.Name `xml:"AIXM-Snapshot"`
	// Attributs de la racine
	Created   string `xml:"created,attr"`
	Origin    string `xml:"origin,attr"`
	Version   string `xml:"version,attr"`
	Effective string `xml:"effective,attr"`
	// Sections principales
	Ahp []Ahp `xml:"Ahp"` // AerodromeHeliportType
	Aha []Aha `xml:"Aha"` // AerodromeHeliportAddressType
	Aho []Aho `xml:"Aho"` // AerodromeHeliportObstacleType
	Agl []Agl `xml:"Agl"` // AeronauticalGroundLightType
	Ahs []Ahs `xml:"Ahs"` // GroundServiceType
	Ahu []Ahu `xml:"Ahu"` // AerodromeHeliportUsageType
	Sah []Sah `xml:"Sah"` // AerodromeHeliportServiceType
	Twy []Twy `xml:"Twy"` // TaxiwayType
	Swy []Swy `xml:"Swy"` // StopwayType
	Apn []Apn `xml:"Apn"` // ApronType
	Gsd []Gsd `xml:"Gsd"` // GateStandType
	Pfy []Pfy `xml:"Pfy"` // PassengerFacilityType
	Rwy []Rwy `xml:"Rwy"` // RunwayType
	Rcp []Rcp `xml:"Rcp"` // RunwayCentreLinePositionType
	Rda []Rda `xml:"Rda"` // RunwayDirectionApproachLightingSystemType
	Rdd []Rdd `xml:"Rdd"` // RunwayDirectionDeclaredDistanceType
	Rdn []Rdn `xml:"Rdn"` // RunwayDirectionType
	Rls []Rls `xml:"Rls"` // RunwayDirectionLightingSystemType
	Rpa []Rpa `xml:"Rpa"` // RunwayProtectionAreaType
	Tla []Tla `xml:"Tla"` // TlofType
	Nsc []Nsc `xml:"Nsc"` // CheckpointType
	Ase []Ase `xml:"Ase"` // AirspaceType
	Abd []Abd `xml:"Abd"` // AirspaceBorderType
	Adg []Adg `xml:"Adg"` // AirspaceDerivedGeometryType
	Fto []Fto `xml:"Fto"` // FatoType
	Gbr []Gbr `xml:"Gbr"` // GeographicalBorderType
	Sae []Sae `xml:"Sae"` // AirspaceServiceType
	Dpn []Dpn `xml:"Dpn"` // DesignatedPointType
	Spa []Spa `xml:"Spa"` // SignificantPointAirspaceType
	Vor []Vor `xml:"Vor"` // VorType
	Ndb []Ndb `xml:"Ndb"` // NdbType
	Spd []Spd `xml:"Spd"` // SpecialDateType
	Dme []Dme `xml:"Dme"` // DmeType
	Tcn []Tcn `xml:"Tcn"` // TcnType
	Ils []Ils `xml:"Ils"` // IlsType
	Mkr []Mkr `xml:"Mkr"` // MkrType
	Rte []Rte `xml:"Rte"` // RouteType
	Rsg []Rsg `xml:"Rsg"` // RouteSegmentType
	Rsu []Rsu `xml:"Rsu"` // RouteSegmentUsageType
	Plb []Plb `xml:"Plb"` // CruisingLevelsTableType
	Plc []Plc `xml:"Plc"` // CruisingLevelsColumnType
	Pll []Pll `xml:"Pll"` // CruisingLevelsList
	Org []Org `xml:"Org"` // OrganisationAuthorityType
	Uni []Uni `xml:"Uni"` // UnitType
	Ser []Ser `xml:"Ser"` // ServiceType
	Fqy []Fqy `xml:"Fqy"` // FrequencyType
	Obs []Obs `xml:"Obs"` // ObstacleType
	Agt []Agt `xml:"Agt"` // AerodromeHeliportAdditionalType
	Ast []Ast `xml:"Ast"` // AerodromeServiceType
	Att []Att `xml:"Att"` // AttributeType
	Avx []Avx `xml:"Avx"` // AirspaceVertexType
	Cdl []Cdl `xml:"Cdl"` // CodeLangType
	Dtt []Dtt `xml:"Dtt"` // DmeTimeType
	Ftt []Ftt `xml:"Ftt"` // FrequencyTimeType
	Gbv []Gbv `xml:"Gbv"` // GeometricBufferVertexType
	Igp []Igp `xml:"Igp"` // IlsGlidePathType
	Igt []Igt `xml:"Igt"` // IlsGlidePathTimeType
	Ilt []Ilt `xml:"Ilt"` // IlsLocalizerType
	Ilz []Ilz `xml:"Ilz"` // IlsLocalizerTimeType
	Mkt []Mkt `xml:"Mkt"` // Temps Balises d'Approche
	Ntt []Ntt `xml:"Ntt"` // NdbTimeType
	Rst []Rst `xml:"Rst"` // RouteSegmentTimingType
	Rul []Rul `xml:"Rul"` // RouteUseDefinitionType
	Stt []Stt `xml:"Stt"` // SignificantPointTimeType
	Ttt []Ttt `xml:"Ttt"` // TacanTimeType
	Vtt []Vtt `xml:"Vtt"` // VorTimeType
}
