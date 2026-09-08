package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/service"
)

func main() {
	// Arguments en ligne de commande
	xmlFile := flag.String("input", "", "Chemin vers le fichier XML AIXM (requis)")
	outputDir := flag.String("output", "", "Répertoire de sortie pour les fichiers JSON (requis)")
	force := flag.Bool("force", false, "Force la conversion même si .last_conversion existe")
	flag.Parse()

	// Vérifier les arguments
	if *xmlFile == "" || *outputDir == "" {
		log.Println("Usage: xml2json -input <fichier.xml> -output <répertoire> [-force]")
		log.Println("\nExemple:")
		log.Println("  xml2json -input db/sia/export_xml_bd_sia_2025-12-25-v02/AIXM4.5_all_FR_OM_2025-12-25.xml -output db/json")
		log.Println("  xml2json -input ... -output ... -force  # Force la reconversion")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Vérifier que le fichier XML existe
	if _, err := os.Stat(*xmlFile); os.IsNotExist(err) {
		log.Fatalf("Erreur: le fichier XML n'existe pas: %s", *xmlFile)
	}

	// Convertir en chemins absolus
	absXMLFile, err := filepath.Abs(*xmlFile)
	if err != nil {
		log.Fatalf("Erreur lors de la résolution du chemin XML: %v", err)
	}

	absOutputDir, err := filepath.Abs(*outputDir)
	if err != nil {
		log.Fatalf("Erreur lors de la résolution du chemin de sortie: %v", err)
	}

	// Si -force, supprimer le fichier .last_conversion
	if *force {
		markerPath := filepath.Join(absOutputDir, ".last_conversion")
		if err := os.Remove(markerPath); err != nil && !os.IsNotExist(err) {
			log.Printf("Avertissement: impossible de supprimer %s: %v", markerPath, err)
		} else if err == nil {
			log.Println("Fichier .last_conversion supprimé (mode force)")
		}
	}

	log.Println("========================================")
	log.Println("  Conversion AIXM XML vers JSON")
	log.Println("========================================")
	log.Printf("Fichier d'entrée: %s", absXMLFile)
	log.Printf("Répertoire de sortie: %s", absOutputDir)
	log.Println("========================================")

	// Lancer le processus de conversion
	if err := service.ProcessAIXMToJSON(absXMLFile, absOutputDir); err != nil {
		log.Fatalf("Erreur lors de la conversion: %v", err)
	}

	log.Println("========================================")
	log.Println("  Conversion terminée avec succès!")
	log.Println("========================================")
}
