package json

type ADSBResponse struct {
	Ac []ADSBAircraft `json:"ac"`
}

type ADSBAircraft struct {
	Flight string      `json:"flight"`
	Lat    float64     `json:"lat"`
	Lon    float64     `json:"lon"`
	Alt    interface{} `json:"alt_baro"`
	Track  float64     `json:"track"`
	GS     float64     `json:"gs"`
	Type   string      `json:"t"`
}

type VRSRoute struct {
	Route []VRSRoutePoint `json:"_route"`
}

type VRSRoutePoint struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
}
