package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// AirplaneFetcher gère la récupération des données d'avions depuis une API externe
type AirplaneFetcher struct {
	client  *http.Client
	baseURL string
	timeout time.Duration
}

// NewAirplaneFetcher crée une nouvelle instance du fetcher
func NewAirplaneFetcher(baseURL string, timeout time.Duration) *AirplaneFetcher {
	return &AirplaneFetcher{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
		timeout: timeout,
	}
}

// FetchAirplanes récupère la liste des avions depuis l'API
func (af *AirplaneFetcher) FetchAirplanes(apiURL string) (map[string]interface{}, error) {
	resp, err := af.client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'appel API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("réponse API inattendue (%d): %s", resp.StatusCode, string(body))
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture de la réponse: %w", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON: %w", err)
	}

	return payload, nil
}

// FetchRoute récupère la route d'un avion depuis l'API de standing data
func (af *AirplaneFetcher) FetchRoute(prefix, callsign string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://vrs-standing-data.adsb.lol/routes/%s/%s.json", prefix, callsign)

	resp, err := af.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération de la route: %w", err)
	}
	defer resp.Body.Close()

	// Si la route n'existe pas, retourner nil sans erreur
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("réponse API inattendue (%d): %s", resp.StatusCode, string(body))
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture de la réponse: %w", err)
	}

	var route map[string]interface{}
	if err := json.Unmarshal(raw, &route); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON de la route: %w", err)
	}

	return route, nil
}
