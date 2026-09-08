package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/engine"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/handler"
	"github.com/vdessenn/air-traffic-simulator/backend/internal/service"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("=== AI30 AirTraffic Simulation Starting ===")

	// Déterminer les chemins des fichiers
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du répertoire courant: %v", err)
	}

	// Déterminer la racine du projet
	// Si on est dans le répertoire backend/, remonter d'un niveau
	projectRoot := workDir
	if filepath.Base(workDir) == "backend" {
		projectRoot = filepath.Dir(workDir)
	}

	// Chercher le fichier AIXM4.5 le plus récent
	xmlPath, err := service.FindLatestAIXMFile(projectRoot)
	if err != nil {
		log.Fatalf("Erreur lors de la recherche du fichier AIXM: %v", err)
	}

	jsonOutputDir := filepath.Join(projectRoot, "db/json")

	// Vérifier et convertir les données si nécessaire (max 28 jours)
	log.Println("=== Vérification des données AIXM ===")
	if err := service.EnsureJSONData(xmlPath, jsonOutputDir, 28); err != nil {
		log.Fatalf("Erreur lors de la gestion des données: %v", err)
	}

	// Initialize AIXM data loader
	dataLoader := service.NewDataLoader(xmlPath)

	// Load AIXM data at startup
	if err := dataLoader.LoadData(); err != nil {
		log.Fatalf("Erreur lors du chargement des données AIXM: %v", err)
	}

	snapshot := dataLoader.GetSnapshot()
	log.Printf("Simulation initialisée avec %d aérodromes et %d espaces aériens",
		len(snapshot.Ahp), len(snapshot.Ase))

	// Initialize simulation engine
	sim := engine.NewSimulation()
	sim.InitEnvironment(dataLoader)

	snapshotPath := filepath.Join(projectRoot, "db/json/planes_snapshot.json")
	routesDir := filepath.Join(projectRoot, "db/routes")

	// Check if test data path is provided via environment variable
	testDataPath := os.Getenv("TEST_DATA_PATH")
	if testDataPath != "" {
		log.Printf("Test mode: Loading data from %s", testDataPath)
		if err := sim.LoadStraightLineTestData(testDataPath); err != nil {
			log.Printf("Erreur lors du chargement des données de test: %v", err)
			log.Printf("Fallback: génération de trafic de test")
			sim.LoadTestData()
		}
		sim.RegisterAllAircraftAgents()
	} else {
		// Normal mode: load real traffic or generate test data
		// Vérifier et télécharger les routes si nécessaire (max 7 jours)
		log.Println("=== Vérification des routes des avions ===")
		if err := service.EnsureRoutesData(snapshotPath, routesDir, 7); err != nil {
			log.Printf("Avertissement lors de la gestion des routes: %v", err)
			// Ne pas bloquer le démarrage en cas d'erreur
		}

		if _, err := os.Stat(snapshotPath); err == nil {
			if err := sim.LoadRealTraffic(snapshotPath, routesDir); err != nil {
				log.Printf("Erreur : %v", err)
			} else {
				sim.RegisterAllAircraftAgents()
			}
		} else {
			// Fallback: no snapshot available, generate test traffic
			log.Printf("Aucun snapshot trouvé, génération de trafic de test")
			sim.LoadTestData()
			sim.RegisterAllAircraftAgents()
		}
	}

	//sim.LoadTestData2()
	//     log.Println("Génération du trafic réaliste aléatoirement...")
	//     sim.SpawnRandomTraffic(100)

	// Start simulation in background
	go sim.Start()

	// Setup handlers
	planeHandler := handler.NewPlaneHandler(sim)
	waypointHandler := handler.NewWaypointHandler(sim)
	riskHandler := handler.NewRiskHandler(sim)
	sectorHandler := handler.NewSectorHandler(sim)
	settingsHandler := handler.NewSettingsHandler(sim)

	// Setup HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/planes", planeHandler.GetPlanes)
	mux.HandleFunc("/tags", waypointHandler.GetTags)
	mux.HandleFunc("/risks", riskHandler.GetRisks)
	mux.HandleFunc("/risks/summary", riskHandler.GetRisksSummary)
	mux.HandleFunc("/sectors", sectorHandler.GetSectors)
	mux.HandleFunc("/simulationSpeed", settingsHandler.SetSpeedFactor)

	// Add health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","airports":` + fmt.Sprintf("%d", len(snapshot.Ahp)) + `,"airspaces":` + fmt.Sprintf("%d", len(snapshot.Ase)) + `}`))
	})

	// Check if React development server is running (dev mode) or serve static files (prod mode)
	reactURL := os.Getenv("REACT_URL")
	var proxyHandler http.Handler
	var staticHandler http.Handler

	if reactURL != "" {
		// Development mode: proxy to React dev server
		if parsedURL, err := url.Parse(reactURL); err == nil {
			proxyHandler = httputil.NewSingleHostReverseProxy(parsedURL)
			log.Printf("Mode développement: Proxy vers React dev server à %s", reactURL)
		} else {
			log.Printf("Avertissement: REACT_URL invalide: %s", reactURL)
		}
	} else {
		// Production mode: serve static files from build directory
		buildDir := filepath.Join(projectRoot, "frontend/simulation-front-end/build")
		if stat, err := os.Stat(buildDir); err == nil && stat.IsDir() {
			staticHandler = http.FileServer(http.Dir(buildDir))
			log.Printf("Mode production: Serving static files from %s", buildDir)
		} else {
			log.Printf("Avertissement: Build directory not found at %s", buildDir)
			log.Printf("Run 'make prod' to build the frontend")
		}
	}

	// Create wrapper handler that routes requests appropriately
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// API routes take precedence
		if strings.HasPrefix(r.URL.Path, "/planes") ||
			strings.HasPrefix(r.URL.Path, "/tags") ||
			strings.HasPrefix(r.URL.Path, "/risks") ||
			strings.HasPrefix(r.URL.Path, "/risks/summary") ||
			strings.HasPrefix(r.URL.Path, "/sectors") ||
			strings.HasPrefix(r.URL.Path, "/simulationSpeed") ||
			strings.HasPrefix(r.URL.Path, "/health") ||
			strings.HasPrefix(r.URL.Path, "/data/") {
			mux.ServeHTTP(w, r)
			return
		}

		// Development mode: proxy to React dev server
		if proxyHandler != nil {
			proxyHandler.ServeHTTP(w, r)
			return
		}

		// Production mode: serve static files
		if staticHandler != nil {
			// For SPA routing: serve index.html for non-file requests
			path := filepath.Join(projectRoot, "frontend/simulation-front-end/build", r.URL.Path)
			if _, err := os.Stat(path); os.IsNotExist(err) && !strings.Contains(r.URL.Path, ".") {
				// File doesn't exist and no extension = SPA route, serve index.html
				r.URL.Path = "/"
			}

			// Add cache control headers
			if strings.HasSuffix(r.URL.Path, ".html") || r.URL.Path == "/" {
				// No cache for HTML files (to avoid stale index.html)
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				w.Header().Set("Pragma", "no-cache")
				w.Header().Set("Expires", "0")
			} else if strings.Contains(r.URL.Path, "/static/") {
				// Long cache for hashed static assets
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			}

			staticHandler.ServeHTTP(w, r)
			return
		}

		// Fallback: API info (if no static files available)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"AI30 AirTraffic API","version":"1.0.0","endpoints":["/","/health","/planes","/tags","/data/","/risks"]}`))
	})

	mux.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir(jsonOutputDir))))

	corsHandler := corsMiddleware(finalHandler)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "7500"
	}

	log.Printf("Serveur démarreur sur le port %s...", port)
	if err := http.ListenAndServe(":"+port, corsHandler); err != nil {
		log.Fatal(err)
	}
}
