package external

import (
	"Backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Quote struct {
	Current       float64 `json:"c"`  // Prezzo attuale
	Change        float64 `json:"d"`  // Variazione assoluta
	Percent       float64 `json:"dp"` // Variazione percentuale
	High          float64 `json:"h"`
	Low           float64 `json:"l"`
	Open          float64 `json:"o"`
	PreviousClose float64 `json:"pc"`
	Timestamp     int64   `json:"t"` // UNIX seconds
}

func FetchLastQuote(ticker string) (models.Price_Change, error) {
	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		return models.Price_Change{}, fmt.Errorf("FINNHUB_API_KEY non trovata")
	}

	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", ticker, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return models.Price_Change{}, fmt.Errorf("errore durante la richiesta: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Price_Change{}, fmt.Errorf("errore HTTP: %s", resp.Status)
	}

	var respQuote Quote
	if err := json.NewDecoder(resp.Body).Decode(&respQuote); err != nil {
		return models.Price_Change{}, fmt.Errorf("errore nel decoding JSON: %w", err)
	}

	priceChange := models.Price_Change{
		Current_price: respQuote.Current,
		DailyChange:   respQuote.Percent,
	}

	return priceChange, nil
}
