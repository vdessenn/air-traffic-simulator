package service

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/vdessenn/air-traffic-simulator/backend/internal/infrastructure/persistence/aixm"
)

// DataLoader gère le chargement séquentiel des données XML
type DataLoader struct {
	xmlFilePath string
	snapshot    *aixm.AIXMSnapshot
}

// NewDataLoader crée une nouvelle instance de DataLoader
func NewDataLoader(xmlFilePath string) *DataLoader {
	return &DataLoader{
		xmlFilePath: xmlFilePath,
		snapshot:    &aixm.AIXMSnapshot{},
	}
}

// LoadData charge le fichier XML de manière séquentielle
func (dl *DataLoader) LoadData() error {
	log.Printf("Début du chargement du fichier XML: %s", dl.xmlFilePath)

	file, err := os.Open(dl.xmlFilePath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier XML: %w", err)
	}
	defer file.Close()

	// Utiliser xml.Decoder pour un parsing séquentiel
	decoder := xml.NewDecoder(file)

	// Parser le fichier XML complet dans la structure AIXMSnapshot
	if err := decoder.Decode(dl.snapshot); err != nil && err != io.EOF {
		return fmt.Errorf("erreur lors du décodage XML: %w", err)
	}

	log.Printf("Chargement XML terminé avec succès")
	dl.logStatistics()

	return nil
}

// logStatistics affiche les statistiques de chargement
func (dl *DataLoader) logStatistics() {
	log.Println("================== Statistiques de chargement AIXM ==================")

	// Catégorie Aérodromes et Pistes
	log.Println("\n[AERODROMES ET PISTES]")
	log.Printf("  Aérodromes/Héliports (Ahp): %d", len(dl.snapshot.Ahp))
	log.Printf("  Adresses Aérodrome (Aha): %d", len(dl.snapshot.Aha))
	log.Printf("  Obstacles Aérodrome (Aho): %d", len(dl.snapshot.Aho))
	log.Printf("  Lumières au Sol (Agl): %d", len(dl.snapshot.Agl))
	log.Printf("  Services au Sol (Ahs): %d", len(dl.snapshot.Ahs))
	log.Printf("  Utilisation Aérodrome (Ahu): %d", len(dl.snapshot.Ahu))
	log.Printf("  Services Aérodrome (Sah): %d", len(dl.snapshot.Sah))
	log.Printf("  Pistes (Rwy): %d", len(dl.snapshot.Rwy))
	log.Printf("  Axe Piste (Rcp): %d", len(dl.snapshot.Rcp))
	log.Printf("  Éclairage Approche Piste (Rda): %d", len(dl.snapshot.Rda))
	log.Printf("  Distances Déclarées Piste (Rdd): %d", len(dl.snapshot.Rdd))
	log.Printf("  Désignation Piste (Rdn): %d", len(dl.snapshot.Rdn))
	log.Printf("  Éclairage Piste (Rls): %d", len(dl.snapshot.Rls))
	log.Printf("  Zone Protection Piste (Rpa): %d", len(dl.snapshot.Rpa))
	log.Printf("  Voies de Circulation (Twy): %d", len(dl.snapshot.Twy))
	log.Printf("  Voies d'Arrêt (Swy): %d", len(dl.snapshot.Swy))
	log.Printf("  Rampes/Aprons (Apn): %d", len(dl.snapshot.Apn))
	log.Printf("  Portes Embarquement (Gsd): %d", len(dl.snapshot.Gsd))
	log.Printf("  Installations Passagers (Pfy): %d", len(dl.snapshot.Pfy))
	log.Printf("  Pistes Hélicoptère TLOF (Tla): %d", len(dl.snapshot.Tla))
	log.Printf("  FATO (Fto): %d", len(dl.snapshot.Fto))
	log.Printf("  Checkpoints (Nsc): %d", len(dl.snapshot.Nsc))

	// Catégorie Espaces Aériens
	log.Println("\n[ESPACES AERIENS]")
	log.Printf("  Espaces Aériens (Ase): %d", len(dl.snapshot.Ase))
	log.Printf("  Bordures Espaces Aériens (Abd): %d", len(dl.snapshot.Abd))
	log.Printf("  Géométrie Dérivée (Adg): %d", len(dl.snapshot.Adg))
	log.Printf("  Frontières Géographiques (Gbr): %d", len(dl.snapshot.Gbr))
	log.Printf("  Services Espace Aérien (Sae): %d", len(dl.snapshot.Sae))
	log.Printf("  Points Significatifs Espace (Spa): %d", len(dl.snapshot.Spa))
	log.Printf("  Vertex Espace Aérien (Avx): %d", len(dl.snapshot.Avx))
	log.Printf("  Vertex Buffer Géométrique (Gbv): %d", len(dl.snapshot.Gbv))

	// Catégorie Navigation (Navaids)
	log.Println("\n[AIDES A LA NAVIGATION]")
	log.Printf("  VOR: %d", len(dl.snapshot.Vor))
	log.Printf("  NDB: %d", len(dl.snapshot.Ndb))
	log.Printf("  DME: %d", len(dl.snapshot.Dme))
	log.Printf("  TACAN (Tcn): %d", len(dl.snapshot.Tcn))
	log.Printf("  ILS: %d", len(dl.snapshot.Ils))
	log.Printf("  ILS Glide Path (Igp): %d", len(dl.snapshot.Igp))
	log.Printf("  ILS Localizer (Ilt): %d", len(dl.snapshot.Ilt))
	log.Printf("  Balises Approche (Mkr): %d", len(dl.snapshot.Mkr))

	// Catégorie Points Significatifs
	log.Println("\n[POINTS SIGNIFICATIFS]")
	log.Printf("  Points Significatifs (Dpn): %d", len(dl.snapshot.Dpn))

	// Catégorie Routes
	log.Println("\n[ROUTES]")
	log.Printf("  Routes (Rte): %d", len(dl.snapshot.Rte))
	log.Printf("  Segments de Route (Rsg): %d", len(dl.snapshot.Rsg))
	log.Printf("  Utilisation Segments Route (Rsu): %d", len(dl.snapshot.Rsu))
	log.Printf("  Tableau Niveaux Croisière (Plb): %d", len(dl.snapshot.Plb))
	log.Printf("  Colonne Niveaux Croisière (Plc): %d", len(dl.snapshot.Plc))
	log.Printf("  Liste Niveaux Croisière (Pll): %d", len(dl.snapshot.Pll))
	log.Printf("  Définitions Usage Route (Rul): %d", len(dl.snapshot.Rul))

	// Catégorie Organisations
	log.Println("\n[ORGANISATIONS ET SERVICES]")
	log.Printf("  Organisations (Org): %d", len(dl.snapshot.Org))
	log.Printf("  Unités (Uni): %d", len(dl.snapshot.Uni))
	log.Printf("  Services (Ser): %d", len(dl.snapshot.Ser))
	log.Printf("  Fréquences (Fqy): %d", len(dl.snapshot.Fqy))

	// Catégorie Obstacles
	log.Println("\n[OBSTACLES]")
	log.Printf("  Obstacles (Obs): %d", len(dl.snapshot.Obs))

	// Catégorie Temps
	log.Println("\n[DONNEES TEMPORELLES]")
	log.Printf("  Dates Spéciales (Spd): %d", len(dl.snapshot.Spd))
	log.Printf("  DME Time (Dtt): %d", len(dl.snapshot.Dtt))
	log.Printf("  Frequency Time (Ftt): %d", len(dl.snapshot.Ftt))
	log.Printf("  ILS Glide Path Time (Igt): %d", len(dl.snapshot.Igt))
	log.Printf("  ILS Localizer Time (Ilz): %d", len(dl.snapshot.Ilz))
	log.Printf("  Marker Time (Mkt): %d", len(dl.snapshot.Mkt))
	log.Printf("  NDB Time (Ntt): %d", len(dl.snapshot.Ntt))
	log.Printf("  Route Segment Timing (Rst): %d", len(dl.snapshot.Rst))
	log.Printf("  Significant Point Time (Stt): %d", len(dl.snapshot.Stt))
	log.Printf("  TACAN Time (Ttt): %d", len(dl.snapshot.Ttt))
	log.Printf("  VOR Time (Vtt): %d", len(dl.snapshot.Vtt))

	// Catégorie Attributs
	log.Println("\n[ATTRIBUTS]")
	log.Printf("  Attributs (Att): %d", len(dl.snapshot.Att))
	log.Printf("  Code Language (Cdl): %d", len(dl.snapshot.Cdl))

	log.Println("\n=====================================================================")
}

// GetSnapshot retourne le snapshot AIXM chargé
func (dl *DataLoader) GetSnapshot() *aixm.AIXMSnapshot {
	return dl.snapshot
}
