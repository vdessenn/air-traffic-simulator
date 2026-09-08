package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// RoutesDownloader gère le téléchargement des routes depuis VRS Standing Data
type RoutesDownloader struct {
	baseURL    string
	client     *http.Client
	timeout    time.Duration
	limiter    *rate.Limiter
	maxRetries int
}

// NewRoutesDownloader crée une nouvelle instance du téléchargeur de routes
func NewRoutesDownloader(timeout time.Duration) *RoutesDownloader {

	// Augmenter le taux de requêtes: 20 req/s avec burst de 10
	limiter := rate.NewLimiter(rate.Limit(20), 10)

	return &RoutesDownloader{
		baseURL: "https://vrs-standing-data.adsb.lol/routes",
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxConnsPerHost:       200,              // Augmenté significativement pour haute concurrence
				MaxIdleConnsPerHost:   100,              // Plus de connexions idle par host
				MaxIdleConns:          500,              // Pool global massif
				IdleConnTimeout:       90 * time.Second, // Réduit légèrement pour libérer plus vite
				DisableKeepAlives:     false,            // Keep-alive crucial pour la performance
				DisableCompression:    false,
				TLSHandshakeTimeout:   5 * time.Second, // Réduit pour timeout plus rapide
				ResponseHeaderTimeout: 5 * time.Second, // Réduit pour détecter les problèmes plus vite
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
		timeout:    timeout,
		limiter:    limiter,
		maxRetries: 1, // Un seul retry pour maximiser la vitesse
	}
}

// DownloadRoutesForCallsigns télécharge les routes pour une liste de callsigns spécifiques
func (rd *RoutesDownloader) DownloadRoutesForCallsigns(callsigns []string, outputDir string) error {
	log.Printf("Téléchargement des routes pour %d callsigns", len(callsigns))

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("erreur lors de la création du répertoire %s: %w", outputDir, err)
	}

	var wg sync.WaitGroup
	errorChan := make(chan error, len(callsigns))
	successCount := 0
	mu := sync.Mutex{}

	// Limiter la concurrence pour éviter de surcharger l'API
	semaphore := make(chan struct{}, 200) // 100 téléchargements parallèles max

	for _, callsign := range callsigns {
		wg.Add(1)
		go func(cs string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			cleanCS := strings.TrimSpace(cs)
			if len(cleanCS) < 3 {
				return
			}

			prefix := cleanCS[:2]
			if err := rd.downloadRouteFile(prefix, cleanCS, outputDir); err != nil {
				if !strings.Contains(err.Error(), "réponse API inattendue pour") {
					errorChan <- err
				}
			} else {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}(callsign)
	}

	wg.Wait()
	close(errorChan)

	// Afficher les erreurs
	errorCount := 0
	for err := range errorChan {
		log.Printf("Erreur: %v", err)
		errorCount++
	}

	log.Printf("Téléchargement terminé: %d/%d callsigns traités avec succès", successCount, len(callsigns))
	if errorCount > 0 {
		log.Printf("%d erreurs rencontrées", errorCount)
	}

	return nil
}

// downloadRouteFile télécharge un fichier de route spécifique
func (rd *RoutesDownloader) downloadRouteFile(prefix, callsign, outputDir string) error {
	var lastErr error

	for attempt := 1; attempt <= rd.maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), rd.timeout)
		if err := rd.limiter.Wait(ctx); err != nil {
			cancel()
			return fmt.Errorf("erreur de limitation de débit pour %s: %w", callsign, err)
		}
		cancel()

		url := fmt.Sprintf("%s/%s/%s.json", rd.baseURL, prefix, callsign)

		// Log pour les nouvelles tentatives
		if attempt > 1 {
			log.Printf("Nouvelle tentative %d/%d pour %s", attempt, rd.maxRetries, callsign)
		}

		resp, err := rd.client.Get(url)
		if err != nil {
			lastErr = err

			if isEnhanceYourCalm(err) {
				waitTime := time.Duration(1<<uint(attempt-1)) * time.Second
				log.Printf("Limite de débit atteinte pour %s, attente de %v avant la prochaine tentative", callsign, waitTime)
				time.Sleep(waitTime)
				continue
			}

			// Pour d'autres erreurs, on réessaie jusqu'à maxRetries
			if attempt < rd.maxRetries {
				time.Sleep(time.Duration(attempt) * 200 * time.Millisecond) // Réduit de 500ms à 200ms
				continue
			}

			return fmt.Errorf("erreur lors du téléchargement de %s: %w", callsign, err)
		}

		err = rd.processResponse(resp, prefix, callsign, outputDir)
		defer resp.Body.Close()

		if err != nil {
			if resp.StatusCode == http.StatusNotFound {
				return nil
			}

			lastErr = err

			// Retry seulement pour les erreurs réseau, pas pour les erreurs de parsing
			if attempt < rd.maxRetries && isNetworkError(err) {
				time.Sleep(time.Duration(attempt) * 200 * time.Millisecond) // Réduit de 500ms à 200ms
				continue
			}

			return err
		}

		return nil
	}
	return lastErr
}

// processResponse traite la réponse HTTP et enregistre le fichier
func (rd *RoutesDownloader) processResponse(resp *http.Response, prefix, callsign, outputDir string) error {

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("réponse API inattendue pour %s (%d)", callsign, resp.StatusCode)
	}

	// Créer le répertoire du préfixe
	prefixDir := filepath.Join(outputDir, prefix)
	if err := os.MkdirAll(prefixDir, 0755); err != nil {
		return fmt.Errorf("erreur lors de la création du répertoire: %w", err)
	}

	// Lire le contenu
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture de la réponse pour %s: %w", callsign, err)
	}

	// Valider et formater le JSON
	var route map[string]interface{}
	if err := json.Unmarshal(raw, &route); err != nil {
		return fmt.Errorf("erreur lors du parsing JSON pour %s: %w", callsign, err)
	}

	// Écrire le fichier formaté
	formatted, _ := json.MarshalIndent(route, "", "  ")
	filePath := filepath.Join(prefixDir, callsign+".json")

	if err := os.WriteFile(filePath, formatted, 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier %s: %w", filePath, err)
	}

	return nil
}

// isEhnaceYourCalm vérifie si l'erreur est liée à une limitation de débit
func isEnhanceYourCalm(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "GOAWAY") ||
		strings.Contains(errMsg, "ENHANCE_YOUR_CALM") ||
		strings.Contains(errMsg, "server sent GOAWAY")
}

// isNetworkError détecte si l'erreur est une erreur réseau (retry-able)
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	// Erreurs réseau typiques : connection reset, timeout, etc.
	return strings.Contains(errMsg, "connection") ||
		strings.Contains(errMsg, "reset") ||
		strings.Contains(errMsg, "timeout") ||
		strings.Contains(errMsg, "EOF") ||
		strings.Contains(errMsg, "GOAWAY")
}

// ExtractCallsignsFromSnapshot extrait les callsigns du fichier planes_snapshot.json
func ExtractCallsignsFromSnapshot(snapshotPath string) ([]string, error) {
	data, err := os.ReadFile(snapshotPath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du snapshot: %w", err)
	}

	var snapshot map[string]interface{}
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing du snapshot: %w", err)
	}

	callsigns := []string{}
	ac, ok := snapshot["ac"].([]interface{})
	if !ok {
		return callsigns, nil
	}

	for _, plane := range ac {
		planeMap, ok := plane.(map[string]interface{})
		if !ok {
			continue
		}

		flight, ok := planeMap["flight"].(string)
		if ok && flight != "" {
			callsigns = append(callsigns, flight)
		}
	}

	log.Printf("Extrait %d callsigns du snapshot", len(callsigns))
	return callsigns, nil
}
