package model

type Sector struct {
	ID       string
	Name     string
	Boundary []Coordinate
	Counts   int
}

func (s *Sector) IsPointInside(p Coordinate) bool {
	inside := false
	j := len(s.Boundary) - 1
	for i := 0; i < len(s.Boundary); i++ {
		if (s.Boundary[i].Latitude > p.Latitude) != (s.Boundary[j].Latitude > p.Latitude) &&
			p.Longitude < (s.Boundary[j].Longitude-s.Boundary[i].Longitude)*(p.Latitude-s.Boundary[i].Latitude)/(s.Boundary[j].Latitude-s.Boundary[i].Latitude)+s.Boundary[i].Longitude {
			inside = !inside
		}
		j = i
	}
	return inside
}
