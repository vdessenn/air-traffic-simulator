package main

import (
	"flag"
	"log"
	"path/filepath"
	"time"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/service"
)

func main() {
	snapshot := flag.String("snapshot", filepath.Join("..", "db", "json", "planes_snapshot.json"), "Chemin du fichier planes_snapshot.json")
	output := flag.String("output", filepath.Join("..", "db", "routes"), "Repertoire de sortie pour les routes")
	timeout := flag.Duration("timeout", 30*time.Second, "Timeout de la requete HTTP")
	flag.Parse()

	downloader := service.NewRoutesDownloader(*timeout)

	log.Printf("Extraction des callsigns depuis %s", *snapshot)
	callsigns, err := service.ExtractCallsignsFromSnapshot(*snapshot)
	if err != nil {
		log.Fatalf("Erreur lors de l'extraction des callsigns: %v", err)
	}

	if len(callsigns) == 0 {
		log.Fatalf("Aucun callsign trouve dans le snapshot")
	}

	log.Printf("Telechargement de %d routes vers %s", len(callsigns), *output)
	if err := downloader.DownloadRoutesForCallsigns(callsigns, *output); err != nil {
		log.Fatalf("Erreur lors du telechargement des routes: %v", err)
	}

	log.Println("OK - Telechargement des routes termine avec succes")
}
