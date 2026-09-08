package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/agent"
	modelJson "github.com/vdessenn/air-traffic-simulator/backend/internal/infrastructure/persistence/json"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/model"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/model/object"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/service"
)

func (s *Simulation) InitEnvironment(loader *service.DataLoader) {
	s.mu.Lock()
	defer s.mu.Unlock()

	snapshot := loader.GetSnapshot()

	addWaypoint := func(id string, lat, long float64) {
		coord := model.Coordinate{Latitude: lat, Longitude: long}

		s.Waypoints = append(s.Waypoints, object.Waypoint{
			ID:          id,
			Coordinates: coord,
		})
		s.WaypointsMap[id] = coord
	}

	for _, airport := range snapshot.Ahp {
		lat := parseAIXMCoordinate(airport.GeoLat)
		long := parseAIXMCoordinate(airport.GeoLong)
		if lat != 0 && long != 0 {
			addWaypoint(airport.AhpUid.CodeId, lat, long)
		}
	}

	for _, vor := range snapshot.Vor {
		lat := parseAIXMCoordinate(vor.VorUid.GeoLat)
		long := parseAIXMCoordinate(vor.VorUid.GeoLong)
		if lat != 0 && long != 0 {
			addWaypoint(vor.VorUid.CodeId, lat, long)
		}
	}

	for _, dpn := range snapshot.Dpn {
		addWaypoint(dpn.DpnUid.CodeId, parseAIXMCoordinate(dpn.DpnUid.GeoLat), parseAIXMCoordinate(dpn.DpnUid.GeoLong))
	}

	for _, border := range snapshot.Abd {
		sectorID := border.AbdUid.AseUid.CodeId
		poly := make([]model.Coordinate, 0)

		for _, v := range border.Avx {
			lat := parseAIXMCoordinate(v.GeoLat)
			lon := parseAIXMCoordinate(v.GeoLong)
			if lat != 0 && lon != 0 {
				poly = append(poly, model.Coordinate{Latitude: lat, Longitude: lon})
			}
		}

		if len(poly) > 2 {
			s.Sectors = append(s.Sectors, &model.Sector{
				ID:       sectorID,
				Boundary: poly,
			})
		}
	}
}

func (s *Simulation) LoadRealTraffic(snapshotPath, routesDir string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Load real traffic data from: %s", snapshotPath)

	data, err := os.ReadFile(snapshotPath)
	if err != nil {
		return fmt.Errorf("error reading snapshot: %v", err)
	}

	var adsbData modelJson.ADSBResponse
	if err := json.Unmarshal(data, &adsbData); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	loadedCount := 0
	routeCount := 0

	for _, acData := range adsbData.Ac {
		callsign := strings.TrimSpace(acData.Flight)
		if callsign == "" {
			continue
		}

		altitude := 0
		switch v := acData.Alt.(type) {
		case float64:
			altitude = int(v)
		case int:
			altitude = v
		}

		if altitude <= 0 {
			// continue
		}

		plane := agent.NewAircraft(callsign, agent.AircraftType("COMMERCIAL"))
		plane.Position = agent.Coordinate{
			Latitude:  acData.Lat,
			Longitude: acData.Lon,
			Altitude:  altitude,
		}
		plane.Heading = int(acData.Track)
		plane.Speed = acData.GS

		// routesDir/PREFIX/CALLSIGN.json
		prefix := ""
		if len(callsign) >= 2 {
			prefix = callsign[:2]
		}

		routeFile := filepath.Join(routesDir, prefix, callsign+".json")
		if _, err := os.Stat(routeFile); err == nil {
			flightPlan, err := loadFlightPlanFromFile(routeFile)
			if err == nil && len(flightPlan) > 0 {
				plane.FlightPlan = flightPlan

				plane.Navigate()
				routeCount++
			}
		}

		if len(plane.FlightPlan) > 0 {
			s.Aircrafts[callsign] = plane
			loadedCount++
		}
	}

	log.Printf("Trafic réel chargé : %d avions (%d avec plan de vol)", loadedCount, routeCount)
	return nil
}

func parseAIXMCoordinate(coord string) float64 {
	coord = strings.TrimSpace(coord)
	if len(coord) == 0 {
		return 0
	}

	lastChar := coord[len(coord)-1]
	var direction byte

	if lastChar == 'N' || lastChar == 'S' || lastChar == 'E' || lastChar == 'W' {
		direction = lastChar
		coord = coord[:len(coord)-1]
	}

	val, err := strconv.ParseFloat(coord, 64)
	if err == nil {
		if math.Abs(val) <= 180 {
			if direction == 'S' || direction == 'W' {
				return -val
			}
			return val
		}
	}

	parts := strings.Split(coord, ".")
	mainPart := parts[0]
	secondsDecimal := 0.0
	if len(parts) > 1 {
		secondsDecimal, _ = strconv.ParseFloat("0."+parts[1], 64)
	}

	var d, m, s float64
	l := len(mainPart)

	if l >= 5 {
		sVal, _ := strconv.ParseFloat(mainPart[l-2:], 64)
		s = sVal + secondsDecimal
		m, _ = strconv.ParseFloat(mainPart[l-4:l-2], 64)
		d, _ = strconv.ParseFloat(mainPart[:l-4], 64)
	} else {
		return val
	}

	decimal := d + (m / 60.0) + (s / 3600.0)

	if direction == 'S' || direction == 'W' {
		decimal = -decimal
	}

	return decimal
}

func loadFlightPlanFromFile(path string) ([]agent.Waypoint, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var routeData map[string]interface{}
	if err := json.Unmarshal(content, &routeData); err != nil {
		return nil, err
	}

	var flightPlan []agent.Waypoint

	if airportList, ok := routeData["_airports"].([]interface{}); ok {
		for _, p := range airportList {
			apMap, ok := p.(map[string]interface{})
			if !ok {
				continue
			}

			lat, _ := apMap["lat"].(float64)
			lon, _ := apMap["lon"].(float64)

			name := ""
			if v, ok := apMap["iata"].(string); ok {
				name = v
			}
			if name == "" {
				if v, ok := apMap["icao"].(string); ok {
					name = v
				}
			}
			if name == "" {
				name = "APT"
			}

			flightPlan = append(flightPlan, agent.Waypoint{
				Name:       name,
				Coordinate: agent.Coordinate{Latitude: lat, Longitude: lon},
			})
		}
	}

	return flightPlan, nil
}

func (s *Simulation) RegisterAllAircraftAgents() {
	s.mu.Lock()
	planes := make([]*agent.Aircraft, 0, len(s.Aircrafts))
	for _, p := range s.Aircrafts {
		planes = append(planes, p)
	}
	s.mu.Unlock()

	for _, p := range planes {
		s.registerAircraftAgent(p)
	}
}

func (s *Simulation) registerTowerAgent(tower *agent.TowerControl) {
	s.commMu.Lock()
	if _, exists := s.towers[tower.Callsign]; exists {
		s.commMu.Unlock()
		return
	}

	inbox := make(chan agent.Message, 32768)

	s.towers[tower.Callsign] = tower
	s.inboxes[tower.Callsign] = inbox
	s.commMu.Unlock()

	mailbox := agent.Mailbox{
		Inbox:  inbox,
		Outbox: s.messageBus,
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		tower.Run(mailbox)
	}()
}

func (s *Simulation) registerAircraftAgent(aircraft *agent.Aircraft) {
	s.commMu.Lock()
	if _, exists := s.inboxes[aircraft.Callsign]; exists {
		s.commMu.Unlock()
		return
	}

	// Buffer augmenté pour gérer le flux ADS-B avec 120+ avions (30 NM range)
	inbox := make(chan agent.Message, 512)

	s.inboxes[aircraft.Callsign] = inbox
	s.commMu.Unlock()

	aircraft.SpeedRegEnabled = s.EnableSpeedReg

	mailbox := agent.Mailbox{
		Inbox:  inbox,
		Outbox: s.messageBus,
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		aircraft.Start(mailbox, s.commandQueue)
	}()

	s.commMu.RLock()
	if tower, ok := s.towers[defaultTowerCallsign]; ok {
		tower.AssignAircraft(aircraft)
	}
	s.commMu.RUnlock()
}

func (s *Simulation) LoadTestData() {
	s.mu.Lock()

	lfpgPos, exists := s.WaypointsMap["LFPG"]

	if !exists {
		lfpgPos = model.Coordinate{Latitude: 49.0097, Longitude: 2.5479}
	}

	p1 := agent.NewAircraft("AF100", "COMMERCIAL")
	p1.Position = agent.Coordinate{
		Latitude:  lfpgPos.Latitude,
		Longitude: lfpgPos.Longitude,
		Altitude:  30000,
	}
	p1.Heading = 180
	p1.FlightPlan = []agent.Waypoint{{Name: "LFLL"}}
	s.Aircrafts["AF100"] = p1

	p2 := agent.NewAircraft("BA200", "COMMERCIAL")
	p2.Position = agent.Coordinate{
		Latitude:  lfpgPos.Latitude + 0.001,
		Longitude: lfpgPos.Longitude,
		Altitude:  30000,
	}
	p2.Heading = 360
	p2.FlightPlan = []agent.Waypoint{{Name: "EGLL"}}
	s.Aircrafts["BA200"] = p2

	s.mu.Unlock()

	s.registerAircraftAgent(p1)
	s.registerAircraftAgent(p2)
}

// LoadStraightLineTestData loads test data from straight_line_test_data.json
func (s *Simulation) LoadStraightLineTestData(testDataPath string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Loading straight-line test data from: %s", testDataPath)

	data, err := os.ReadFile(testDataPath)
	if err != nil {
		return fmt.Errorf("error reading test data: %v", err)
	}

	var testData struct {
		Description   string `json:"description"`
		TestScenarios []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Aircraft    []struct {
				Callsign string `json:"callsign"`
				Type     string `json:"type"`
				Position struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
					Altitude  int     `json:"altitude"`
				} `json:"position"`
				Heading      int     `json:"heading"`
				Speed        float64 `json:"speed"`
				OptimalSpeed float64 `json:"optimal_speed"`
				FlightPlan   []struct {
					ID        string  `json:"id"`
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"flight_plan"`
			} `json:"aircraft"`
		} `json:"test_scenarios"`
	}

	if err := json.Unmarshal(data, &testData); err != nil {
		return fmt.Errorf("error parsing test data JSON: %v", err)
	}

	loadedCount := 0

	// Load first scenario by default (can be extended to load specific scenarios)
	if len(testData.TestScenarios) > 0 {
		scenario := testData.TestScenarios[0]
		log.Printf("Loading scenario: %s - %s", scenario.Name, scenario.Description)

		for _, acData := range scenario.Aircraft {
			plane := agent.NewAircraft(acData.Callsign, agent.AircraftType(acData.Type))
			plane.Position = agent.Coordinate{
				Latitude:  acData.Position.Latitude,
				Longitude: acData.Position.Longitude,
				Altitude:  acData.Position.Altitude,
			}
			plane.Heading = acData.Heading
			plane.Speed = acData.Speed
			plane.OptimalSpeed = acData.OptimalSpeed

			// Load flight plan
			if len(acData.FlightPlan) > 0 {
				plane.FlightPlan = make([]agent.Waypoint, len(acData.FlightPlan))
				for i, wp := range acData.FlightPlan {
					plane.FlightPlan[i] = agent.Waypoint{
						Name: wp.ID,
						Coordinate: agent.Coordinate{
							Latitude:  wp.Latitude,
							Longitude: wp.Longitude,
						},
					}

					// Add waypoint to simulation
					s.WaypointsMap[wp.ID] = model.Coordinate{
						Latitude:  wp.Latitude,
						Longitude: wp.Longitude,
					}
				}

				plane.Navigate()
			}

			s.Aircrafts[acData.Callsign] = plane
			loadedCount++
		}
	}

	log.Printf("Straight-line test data loaded: %d aircraft", loadedCount)
	return nil
}
