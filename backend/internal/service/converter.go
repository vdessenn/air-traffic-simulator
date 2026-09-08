package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/infrastructure/persistence/aixm"
)

// AIXMToJSON convertit les données AIXM en JSON et les enregistre dans des fichiers séparés
type AIXMToJSON struct {
	outputDir string
	snapshot  *aixm.AIXMSnapshot
}

// NewAIXMToJSON crée une nouvelle instance de convertisseur
func NewAIXMToJSON(outputDir string, snapshot *aixm.AIXMSnapshot) *AIXMToJSON {
	return &AIXMToJSON{
		outputDir: outputDir,
		snapshot:  snapshot,
	}
}

// ConvertAndSave convertit toutes les données AIXM en JSON et les sauvegarde
func (converter *AIXMToJSON) ConvertAndSave() error {
	log.Println("Début de la conversion AIXM vers JSON")

	// Créer le répertoire de sortie s'il n'existe pas
	if err := os.MkdirAll(converter.outputDir, 0755); err != nil {
		return fmt.Errorf("erreur lors de la création du répertoire de sortie: %w", err)
	}

	// Convertir et sauvegarder chaque type de données
	if err := converter.convertAirports(); err != nil {
		return err
	}
	if err := converter.convertRunways(); err != nil {
		return err
	}
	if err := converter.convertAirspaces(); err != nil {
		return err
	}
	// if err := converter.convertNavaids(); err != nil {
	// 	return err
	// }
	if err := converter.convertRoutes(); err != nil {
		return err
	}
	// if err := converter.convertObstacles(); err != nil {
	// 	return err
	// }
	if err := converter.convertOrganisations(); err != nil {
		return err
	}
	if err := converter.convertSignificantPoints(); err != nil {
		return err
	}
	// if err := converter.convertTime(); err != nil {
	// 	return err
	// }

	log.Println("Conversion AIXM vers JSON terminée avec succès")
	return nil
}

// convertAirports convertit les données d'aérodromes
func (converter *AIXMToJSON) convertAirports() error {
	log.Printf("Conversion des aérodromes (%d éléments)", len(converter.snapshot.Ahp))

	data := make(map[string]interface{})
	data["ahp"] = converter.snapshot.Ahp
	data["aha"] = converter.snapshot.Aha
	data["aho"] = converter.snapshot.Aho
	data["agl"] = converter.snapshot.Agl
	data["ahs"] = converter.snapshot.Ahs
	data["ahu"] = converter.snapshot.Ahu
	data["sah"] = converter.snapshot.Sah
	data["twy"] = converter.snapshot.Twy
	data["swy"] = converter.snapshot.Swy
	data["apn"] = converter.snapshot.Apn
	data["gsd"] = converter.snapshot.Gsd
	data["pfy"] = converter.snapshot.Pfy
	data["nsc"] = converter.snapshot.Nsc
	data["agt"] = converter.snapshot.Agt
	data["ast"] = converter.snapshot.Ast

	return converter.saveToFile("airport.json", data)
}

// convertRunways convertit les données de pistes
func (converter *AIXMToJSON) convertRunways() error {
	log.Printf("Conversion des pistes (%d éléments)", len(converter.snapshot.Rwy))

	data := make(map[string]interface{})
	data["rwy"] = converter.snapshot.Rwy
	data["rcp"] = converter.snapshot.Rcp
	data["rda"] = converter.snapshot.Rda
	data["rdd"] = converter.snapshot.Rdd
	data["rdn"] = converter.snapshot.Rdn
	data["rls"] = converter.snapshot.Rls
	data["rpa"] = converter.snapshot.Rpa
	data["tla"] = converter.snapshot.Tla
	data["fto"] = converter.snapshot.Fto

	return converter.saveToFile("runway.json", data)
}

// convertAirspaces convertit les données d'espaces aériens
func (converter *AIXMToJSON) convertAirspaces() error {
	log.Printf("Conversion des espaces aériens (%d éléments)", len(converter.snapshot.Ase))

	data := make(map[string]interface{})
	data["ase"] = converter.snapshot.Ase
	data["abd"] = converter.snapshot.Abd
	data["adg"] = converter.snapshot.Adg
	data["gbr"] = converter.snapshot.Gbr
	data["sae"] = converter.snapshot.Sae
	data["spa"] = converter.snapshot.Spa
	data["avx"] = converter.snapshot.Avx
	data["gbv"] = converter.snapshot.Gbv

	return converter.saveToFile("airspace.json", data)
}

// convertNavaids convertit les données de navigation
// func (converter *AIXMToJSON) convertNavaids() error {
// 	log.Printf("Conversion des aides à la navigation (VOR: %d, NDB: %d, DME: %d, ILS: %d)",
// 		len(converter.snapshot.Vor), len(converter.snapshot.Ndb),
// 		len(converter.snapshot.Dme), len(converter.snapshot.Ils))

// 	data := make(map[string]interface{})
// 	data["vor"] = converter.snapshot.Vor
// 	data["ndb"] = converter.snapshot.Ndb
// 	data["dme"] = converter.snapshot.Dme
// 	data["tcn"] = converter.snapshot.Tcn
// 	data["ils"] = converter.snapshot.Ils
// 	data["mkr"] = converter.snapshot.Mkr
// 	data["igp"] = converter.snapshot.Igp
// 	data["ilt"] = converter.snapshot.Ilt

// 	return converter.saveToFile("navaids.json", data)
// }

// convertRoutes convertit les données de routes
func (converter *AIXMToJSON) convertRoutes() error {
	log.Printf("Conversion des routes (%d éléments)", len(converter.snapshot.Rte))

	data := make(map[string]interface{})
	data["rte"] = converter.snapshot.Rte
	data["rsg"] = converter.snapshot.Rsg
	data["rsu"] = converter.snapshot.Rsu
	data["plb"] = converter.snapshot.Plb
	data["plc"] = converter.snapshot.Plc
	data["pll"] = converter.snapshot.Pll
	data["rul"] = converter.snapshot.Rul

	return converter.saveToFile("routes.json", data)
}

// convertObstacles convertit les données d'obstacles
// func (converter *AIXMToJSON) convertObstacles() error {
// 	log.Printf("Conversion des obstacles (%d éléments)", len(converter.snapshot.Obs))

// 	data := make(map[string]interface{})
// 	data["obs"] = converter.snapshot.Obs

// 	return converter.saveToFile("obstacles.json", data)
// }

// convertOrganisations convertit les données d'organisations
func (converter *AIXMToJSON) convertOrganisations() error {
	log.Printf("Conversion des organisations (%d éléments)", len(converter.snapshot.Org))

	data := make(map[string]interface{})
	data["org"] = converter.snapshot.Org
	data["uni"] = converter.snapshot.Uni
	data["ser"] = converter.snapshot.Ser
	data["fqy"] = converter.snapshot.Fqy

	return converter.saveToFile("organisation.json", data)
}

// convertSignificantPoints convertit les données de points significatifs
func (converter *AIXMToJSON) convertSignificantPoints() error {
	log.Printf("Conversion des points significatifs (%d éléments)", len(converter.snapshot.Dpn))

	data := make(map[string]interface{})
	data["dpn"] = converter.snapshot.Dpn

	return converter.saveToFile("significantpoints.json", data)
}

// convertTime convertit les données temporelles
// func (converter *AIXMToJSON) convertTime() error {
// 	log.Printf("Conversion des données temporelles")

// 	data := make(map[string]interface{})
// 	data["spd"] = converter.snapshot.Spd
// 	data["dtt"] = converter.snapshot.Dtt
// 	data["ftt"] = converter.snapshot.Ftt
// 	data["igt"] = converter.snapshot.Igt
// 	data["ilz"] = converter.snapshot.Ilz
// 	data["mkt"] = converter.snapshot.Mkt
// 	data["ntt"] = converter.snapshot.Ntt
// 	data["rst"] = converter.snapshot.Rst
// 	data["stt"] = converter.snapshot.Stt
// 	data["ttt"] = converter.snapshot.Ttt
// 	data["vtt"] = converter.snapshot.Vtt
// 	data["att"] = converter.snapshot.Att
// 	data["cdl"] = converter.snapshot.Cdl

// 	return converter.saveToFile("time.json", data)
// }

// saveToFile enregistre les données JSON dans un fichier
func (converter *AIXMToJSON) saveToFile(filename string, data interface{}) error {
	filePath := filepath.Join(converter.outputDir, filename)
	log.Printf("Écriture du fichier: %s", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier %s: %w", filename, err)
	}
	defer file.Close()

	// Normaliser les coordonnées geoLat/geoLong (DMS -> DDD) avant l'écriture
	normalized, err := normalizeGeoCoords(data)
	if err != nil {
		return fmt.Errorf("erreur lors de la normalisation des coordonnées: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Format JSON lisible

	if err := encoder.Encode(normalized); err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON pour %s: %w", filename, err)
	}

	log.Printf("Fichier %s créé avec succès", filename)
	return nil
}

// ProcessAIXMToJSON est la fonction principale pour orchestrer la conversion
func ProcessAIXMToJSON(xmlFilePath, outputDir string) error {
	// Charger les données XML
	loader := NewDataLoader(xmlFilePath)
	if err := loader.LoadData(); err != nil {
		return fmt.Errorf("erreur lors du chargement des données: %w", err)
	}

	// Convertir et sauvegarder en JSON
	converter := NewAIXMToJSON(outputDir, loader.GetSnapshot())
	if err := converter.ConvertAndSave(); err != nil {
		return fmt.Errorf("erreur lors de la conversion: %w", err)
	}

	return nil
}

// --- Normalisation des coordonnées ---

// normalizeGeoCoords convertit récursivement tous les champs geoLat/geoLong exprimés en string (DMS ou DDD)
// en float64 (degrés décimaux). Pour éviter d'écrire des réflections complexes sur les structs AIXM,
// on passe par un marshalling intermédiaire vers une structure générique.
func normalizeGeoCoords(data interface{}) (interface{}, error) {
	// Convertir en représentation générique
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var generic interface{}
	if err := json.Unmarshal(raw, &generic); err != nil {
		return nil, err
	}

	transformed := transformNode(generic)
	return transformed, nil
}

func transformNode(node interface{}) interface{} {
	switch v := node.(type) {
	case map[string]interface{}:
		// Traiter chaque clé
		for k, val := range v {
			lowerK := strings.ToLower(k)
			if lowerK == "geolat" || lowerK == "geolong" {
				switch vv := val.(type) {
				case string:
					v[k] = parseAIXMCoordinateString(vv)
				default:
					// Laisser tel quel si déjà numérique
					v[k] = val
				}
			} else {
				v[k] = transformNode(val)
			}
		}
		return v
	case []interface{}:
		for i, item := range v {
			v[i] = transformNode(item)
		}
		return v
	default:
		return node
	}
}

// parseAIXMCoordinateString convertit un texte AIXM en degrés décimaux (float64).
// Gère les formats:
// - DMS concaténé avec direction: "DDMMSSN" / "DDDMMSSW"
// - Décimal avec/ sans direction: "12.3456", "12.3456N"
func parseAIXMCoordinateString(coord string) float64 {
	coord = strings.TrimSpace(coord)
	if coord == "" {
		return 0
	}

	// Déterminer s'il y a une lettre de direction en suffixe
	var direction byte
	last := coord[len(coord)-1]
	if last == 'N' || last == 'S' || last == 'E' || last == 'W' {
		direction = last
		coord = coord[:len(coord)-1]
	}

	// Tenter DMS avec secondes éventuellement décimales: DDMMSS(.s+)? ou DDDMMSS(.s+)?
	// Exemple: 454402.00 (-> 45°44'02.00") ou 0010019.00 (-> 1°00'19.00")
	if m := regexp.MustCompile(`^(\d{2,3})(\d{2})(\d{2}(?:\.\d+)?)$`).FindStringSubmatch(coord); m != nil {
		d, err1 := strconv.ParseFloat(m[1], 64)
		mnt, err2 := strconv.ParseFloat(m[2], 64)
		sec, err3 := strconv.ParseFloat(m[3], 64)
		if err1 == nil && err2 == nil && err3 == nil {
			dec := d + mnt/60 + sec/3600
			if direction == 'S' || direction == 'W' {
				dec = -dec
			}
			return dec
		}
	}

	// Essayer DMS entiers (sans décimales sur les secondes): 6 (DDMMSS) ou 7 (DDDMMSS)
	if len(coord) == 6 || len(coord) == 7 {
		var d, mnt, sec float64
		var err error
		if len(coord) == 6 {
			d, err = strconv.ParseFloat(coord[:2], 64)
			if err != nil {
				return 0
			}
			mnt, err = strconv.ParseFloat(coord[2:4], 64)
			if err != nil {
				return 0
			}
			sec, err = strconv.ParseFloat(coord[4:], 64)
			if err != nil {
				return 0
			}
		} else {
			d, err = strconv.ParseFloat(coord[:3], 64)
			if err != nil {
				return 0
			}
			mnt, err = strconv.ParseFloat(coord[3:5], 64)
			if err != nil {
				return 0
			}
			sec, err = strconv.ParseFloat(coord[5:], 64)
			if err != nil {
				return 0
			}
		}
		dec := d + mnt/60 + sec/3600
		if direction == 'S' || direction == 'W' {
			dec = -dec
		}
		return dec
	}

	// Essayer le décimal simple
	if val, err := strconv.ParseFloat(coord, 64); err == nil {
		if direction == 'S' || direction == 'W' {
			return -val
		}
		return val
	}

	return 0
}
