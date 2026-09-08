package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// EnsureJSONData vérifie si les données JSON existent et sont à jour
// Si nécessaire, il déclenche la conversion XML -> JSON
func EnsureJSONData(xmlFilePath, jsonOutputDir string, maxAgeDays int) error {
	const markerFile = ".last_conversion"
	markerPath := jsonOutputDir + "/" + markerFile

	// Vérifier si conversion est nécessaire
	needsConversion := shouldConvertData(jsonOutputDir, markerPath, maxAgeDays)

	if needsConversion {
		log.Println("=== Conversion des données AIXM vers JSON ===")

		// Charger les données XML
		loader := NewDataLoader(xmlFilePath)
		if err := loader.LoadData(); err != nil {
			return fmt.Errorf("erreur lors du chargement des données XML: %w", err)
		}

		// Convertir et sauvegarder en JSON
		converter := NewAIXMToJSON(jsonOutputDir, loader.GetSnapshot())
		if err := converter.ConvertAndSave(); err != nil {
			return fmt.Errorf("erreur lors de la conversion: %w", err)
		}

		// Créer le fichier marqueur
		if err := os.MkdirAll(jsonOutputDir, 0755); err != nil {
			return fmt.Errorf("erreur lors de la création du répertoire: %w", err)
		}
		if err := os.WriteFile(markerPath, []byte(time.Now().Format(time.RFC3339)), 0644); err != nil {
			return fmt.Errorf("erreur lors de la création du fichier marqueur: %w", err)
		}

		log.Println("=== Conversion terminée avec succès ===")
	} else {
		log.Println("=== Données JSON déjà à jour, conversion ignorée ===")
	}

	return nil
}

// shouldConvertData détermine si la conversion est nécessaire
func shouldConvertData(jsonDir, markerPath string, maxAgeDays int) bool {
	// Vérifier si le répertoire existe et n'est pas vide
	entries, err := os.ReadDir(jsonDir)
	if err != nil || len(entries) == 0 {
		log.Println("Aucune donnée JSON trouvée")
		return true
	}

	// Vérifier si le fichier marqueur existe
	markerInfo, err := os.Stat(markerPath)
	if err != nil {
		log.Println("Fichier marqueur absent")
		return true
	}

	// Vérifier l'âge du fichier marqueur
	age := time.Since(markerInfo.ModTime())
	ageDays := int(age.Hours() / 24)

	if ageDays > maxAgeDays {
		log.Printf("Données trop anciennes (%d jours)", ageDays)
		return true
	}

	log.Printf("Données à jour (%d jours)", ageDays)
	return false
}

// FindLatestAIXMFile cherche le fichier AIXM4.5 XML le plus récent
// dans les dossiers export_xml_bd_sia_* du répertoire SIA
func FindLatestAIXMFile(projectRoot string) (string, error) {
	// Chercher tous les dossiers export_xml_bd_sia_*
	siaDirPattern := filepath.Join(projectRoot, "db/sia/export_xml_bd_sia_*")
	siaDirs, err := filepath.Glob(siaDirPattern)
	if err != nil || len(siaDirs) == 0 {
		return "", fmt.Errorf("aucun dossier SIA trouvé à: %s\n"+
			"Les données AIXM ne sont pas versionnées dans ce dépôt. "+
			"Voir la section \"Data setup\" du README pour les installer.", siaDirPattern)
	}

	// Chercher le fichier XML le plus récent dans les dossiers trouvés
	var latestXML string
	var latestModTime int64

	for _, dir := range siaDirs {
		files, err := filepath.Glob(filepath.Join(dir, "AIXM4.5_all_FR_OM_*.xml"))
		if err != nil {
			continue
		}

		for _, f := range files {
			info, err := os.Stat(f)
			if err != nil {
				continue
			}

			if info.ModTime().Unix() > latestModTime {
				latestModTime = info.ModTime().Unix()
				latestXML = f
			}
		}
	}

	if latestXML == "" {
		return "", fmt.Errorf("aucun fichier AIXM4.5_all_FR_OM_*.xml trouvé dans les dossiers SIA\n" +
			"Voir la section \"Data setup\" du README.")
	}

	log.Printf("Fichier AIXM trouvé: %s (modifié le %s)",
		filepath.Base(latestXML), time.Unix(latestModTime, 0).Format("2006-01-02 15:04:05"))

	return latestXML, nil
}

// EnsureRoutesData vérifie si les routes des avions existent et sont à jour
// Si nécessaire, il déclenche le téléchargement des routes depuis VRS Standing Data
func EnsureRoutesData(snapshotPath, routesDir string, maxAgeDays int) error {
	const markerFile = ".last_routes_update"
	markerPath := filepath.Join(routesDir, markerFile)

	// Vérifier si le snapshot existe
	if _, err := os.Stat(snapshotPath); err != nil {
		log.Printf("Snapshot introuvable (%s), téléchargement des routes ignoré", snapshotPath)
		return nil
	}

	// Vérifier si mise à jour est nécessaire
	needsUpdate := shouldUpdateRoutes(routesDir, markerPath, maxAgeDays)

	if needsUpdate {
		log.Println("=== Téléchargement/Mise à jour des routes des avions ===")

		// Extraire les callsigns du snapshot
		callsigns, err := ExtractCallsignsFromSnapshot(snapshotPath)
		if err != nil {
			return fmt.Errorf("erreur lors de l'extraction des callsigns: %w", err)
		}

		if len(callsigns) == 0 {
			log.Println("Aucun callsign trouvé dans le snapshot")
			return nil
		}

		// Télécharger les routes avec timeout réduit pour réponses plus rapides
		downloader := NewRoutesDownloader(15 * time.Second) // Réduit de 30s à 15s
		if err := downloader.DownloadRoutesForCallsigns(callsigns, routesDir); err != nil {
			log.Printf("Avertissement: Erreur lors du téléchargement des routes: %v", err)
			// Ne pas bloquer le démarrage en cas d'erreur réseau
			return nil
		}

		// Créer le fichier marqueur
		if err := os.MkdirAll(routesDir, 0755); err != nil {
			return fmt.Errorf("erreur lors de la création du répertoire: %w", err)
		}
		if err := os.WriteFile(markerPath, []byte(time.Now().Format(time.RFC3339)), 0644); err != nil {
			return fmt.Errorf("erreur lors de la création du fichier marqueur: %w", err)
		}

		log.Println("=== Téléchargement des routes terminé avec succès ===")
	} else {
		log.Println("=== Routes des avions déjà à jour, téléchargement ignoré ===")
	}

	return nil
}

// shouldUpdateRoutes détermine si la mise à jour des routes est nécessaire
func shouldUpdateRoutes(routesDir, markerPath string, maxAgeDays int) bool {
	// Vérifier si le répertoire existe et n'est pas vide
	entries, err := os.ReadDir(routesDir)
	if err != nil || len(entries) <= 1 { // <= 1 pour ignorer le fichier marqueur seul
		log.Println("Aucune route trouvée")
		return true
	}

	// Vérifier si le fichier marqueur existe
	markerInfo, err := os.Stat(markerPath)
	if err != nil {
		log.Println("Fichier marqueur des routes absent")
		return true
	}

	// Vérifier l'âge du fichier marqueur
	age := time.Since(markerInfo.ModTime())
	ageDays := int(age.Hours() / 24)

	if ageDays > maxAgeDays {
		log.Printf("Routes trop anciennes (%d jours)", ageDays)
		return true
	}

	log.Printf("Routes à jour (%d jours)", ageDays)
	return false
}
