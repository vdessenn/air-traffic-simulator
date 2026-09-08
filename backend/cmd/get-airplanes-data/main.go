package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/service"
)

// CLI utilitaire pour récupérer les avions depuis une API externe et stocker un snapshot JSON.
// Par défaut, pointe sur un centre France (46.5, 2.5) avec un rayon large pour couvrir métropole + Corse.
// Enrichir un snapshot existant : go run cmd/get-airplanes-data/main.go -enrich-routes -input ../db/json/planes_snapshot.json -output ../db/json/planes_enriched.json
func main() {
	apiURL := flag.String("url", "https://api.adsb.lol/v2/lat/46.5/lon/2.5/dist/500", "URL de l'API qui expose les avions")
	input := flag.String("input", "", "Chemin du fichier JSON d'entrée pour enrichissement (laisser vide pour nouvelle récupération)")
	output := flag.String("output", filepath.Join("..", "db", "json", "planes_snapshot.json"), "Chemin du fichier JSON de sortie")
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout de la requête HTTP")
	enrichRoutes := flag.Bool("enrich-routes", false, "Enrichir les données avec les routes depuis VRS Standing Data")
	routesDir := flag.String("routes-dir", filepath.Join("..", "db", "routes"), "Répertoire contenant les fichiers de routes téléchargés")
	flag.Parse()

	fetcher := service.NewAirplaneFetcher("https://api.adsb.lol", *timeout)

	var planes map[string]interface{}
	var err error

	// Si un fichier d'entrée est spécifié, le charger au lieu de faire un appel API
	if *input != "" {
		log.Printf("Chargement du snapshot depuis %s", *input)
		planes, err = loadSnapshot(*input)
		if err != nil {
			log.Fatalf("Erreur lors du chargement du snapshot: %v", err)
		}
	} else {
		log.Printf("Récupération des avions depuis %s", *apiURL)
		planes, err = fetcher.FetchAirplanes(*apiURL)
		if err != nil {
			log.Fatalf("Erreur lors de la récupération des avions: %v", err)
		}
	}

	// Enrichir avec les routes si demandé
	if *enrichRoutes {
		log.Println("Enrichissement des avions avec les données de routes...")
		startTime := time.Now()

		callsigns := extractCallsignsFromPlanes(planes)
		log.Printf("Callsigns uniques trouvés: %d", len(callsigns))

		downloader := service.NewRoutesDownloader(15 * time.Second)
		if err := downloader.DownloadRoutesForCallsigns(callsigns, *routesDir); err != nil {
			log.Printf("Avertissement lors du téléchargement des routes: %v", err)
		}

		log.Println("Chargement des routes depuis les fichiers locaux...")
		routesMap := loadRoutesFromDisk(*routesDir)
		log.Printf("Routes chargées en mémoire: %d", len(routesMap))

		log.Println("Enrichissement in-memory des avions avec routes...")
		planes = enrichAirplanesWithLocalRoutes(planes, routesMap)

		log.Printf("Enrichissement terminé en %v", time.Since(startTime))

		// Créer le fichier marqueur pour éviter le re-téléchargement au démarrage de l'API
		markerPath := filepath.Join(*routesDir, ".routes_marker")
		if err := os.WriteFile(markerPath, []byte(time.Now().Format(time.RFC3339)), 0644); err != nil {
			log.Printf("Avertissement: impossible de créer le fichier marqueur: %v", err)
		} else {
			log.Printf("Fichier marqueur créé: %s", markerPath)
		}
	}

	// Formater et sauvegarder
	formatted, err := json.MarshalIndent(planes, "", "  ")
	if err != nil {
		log.Fatalf("Erreur lors du formatage JSON: %v", err)
	}

	if err := os.MkdirAll(filepath.Dir(*output), 0755); err != nil {
		log.Fatalf("Erreur lors de la création du répertoire de sortie: %v", err)
	}

	if err := os.WriteFile(*output, formatted, 0644); err != nil {
		log.Fatalf("Erreur lors de l'écriture du fichier JSON: %v", err)
	}

	log.Printf("Données avion sauvegardées dans %s", *output)
}

// extractCallsignsFromPlanes extrait tous les callsigns uniques des avions
func extractCallsignsFromPlanes(planes map[string]interface{}) []string {
	ac, ok := planes["ac"].([]interface{})
	if !ok {
		return []string{}
	}

	callsignSet := make(map[string]bool)
	for _, plane := range ac {
		planeMap, ok := plane.(map[string]interface{})
		if !ok {
			continue
		}

		flight, ok := planeMap["flight"].(string)
		if !ok || flight == "" {
			continue
		}

		cleanFlight := strings.TrimSpace(flight)
		if len(cleanFlight) >= 3 {
			callsignSet[cleanFlight] = true
		}
	}

	// Convertir le set en slice
	callsigns := make([]string, 0, len(callsignSet))
	for cs := range callsignSet {
		callsigns = append(callsigns, cs)
	}

	return callsigns
}

// loadSnapshot charge un fichier JSON existant
func loadSnapshot(filepath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var planes map[string]interface{}
	if err := json.Unmarshal(data, &planes); err != nil {
		return nil, err
	}

	return planes, nil
}

// loadRoutesFromDisk charge toutes les routes depuis le répertoire local
// Construit une map[callsign]route pour un accès O(1)
func loadRoutesFromDisk(routesDir string) map[string]map[string]interface{} {
	routesMap := make(map[string]map[string]interface{})

	// Parcourir tous les sous-répertoires (préfixes)
	entries, err := os.ReadDir(routesDir)
	if err != nil {
		log.Printf("Erreur lors de la lecture du répertoire routes: %v", err)
		return routesMap
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		prefixDir := filepath.Join(routesDir, entry.Name())
		routeFiles, err := os.ReadDir(prefixDir)
		if err != nil {
			continue
		}

		// Charger chaque fichier de route
		for _, routeFile := range routeFiles {
			if !strings.HasSuffix(routeFile.Name(), ".json") {
				continue
			}

			routePath := filepath.Join(prefixDir, routeFile.Name())
			data, err := os.ReadFile(routePath)
			if err != nil {
				continue
			}

			var route map[string]interface{}
			if err := json.Unmarshal(data, &route); err != nil {
				continue
			}

			// Extraire le callsign du nom de fichier (sans .json)
			callsign := strings.TrimSuffix(routeFile.Name(), ".json")
			routesMap[callsign] = route
		}
	}

	return routesMap
}

// enrichAirplanesWithLocalRoutes enrichit les avions avec routes depuis la map locale
// Pure in-memory join - no API calls, no disk I/O except initial load
func enrichAirplanesWithLocalRoutes(planes map[string]interface{}, routesMap map[string]map[string]interface{}) map[string]interface{} {
	ac, ok := planes["ac"].([]interface{})
	if !ok {
		log.Println("Avertissement: impossible de trouver la liste des avions")
		return planes
	}

	totalPlanes := len(ac)
	enrichedCount := 0

	// Join en O(n) - simple boucle avec lookup O(1)
	for i, plane := range ac {
		planeMap, ok := plane.(map[string]interface{})
		if !ok {
			continue
		}

		flight, ok := planeMap["flight"].(string)
		if !ok || flight == "" {
			continue
		}

		cleanFlight := strings.TrimSpace(flight)

		// Lookup O(1) dans la map
		if route, exists := routesMap[cleanFlight]; exists {
			planeMap["route"] = route
			ac[i] = planeMap
			enrichedCount++
		}
	}

	log.Printf("Avions enrichis: %d/%d (%.1f%%)", enrichedCount, totalPlanes, float64(enrichedCount)/float64(totalPlanes)*100)
	planes["ac"] = ac
	return planes
}
