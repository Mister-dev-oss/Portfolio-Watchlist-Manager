package external

import (
	"Backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Init() error {
	return godotenv.Load()
}

// Candle rappresenta una singola candela OHLC restituita da Polygon
type Candle struct {
	Timestamp int64   `json:"t"`
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Close     float64 `json:"c"`
	Volume    float64 `json:"v"`
}

// polygonResponse rappresenta la risposta grezza dell'API di Polygon
type polygonResponse struct {
	Results []Candle `json:"results"`
}

// FetchOHLC scarica le candele da Polygon per un ticker e un intervallo di date
func FetchOHLC(ticker string, fromDate, toDate time.Time) ([]models.OHLC, error) {
	apiKey := os.Getenv("POLYGON_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("POLYGON_API_KEY non trovata")
	}

	url := fmt.Sprintf(
		"https://api.polygon.io/v2/aggs/ticker/%s/range/1/day/%s/%s?adjusted=true&sort=asc&limit=5000&apiKey=%s",
		ticker,
		fromDate.Format("2006-01-02"),
		toDate.Format("2006-01-02"),
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("errore durante la richiesta: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("errore HTTP: %s", resp.Status)
	}

	var polyResp polygonResponse
	if err := json.NewDecoder(resp.Body).Decode(&polyResp); err != nil {
		return nil, fmt.Errorf("errore nel decoding JSON: %w", err)
	}

	ohlcs := make([]models.OHLC, 0, len(polyResp.Results))
	for _, c := range polyResp.Results {
		dateStr := time.UnixMilli(c.Timestamp).Format("2006-01-02")
		ohlcs = append(ohlcs, models.OHLC{
			Date:   dateStr,
			Open:   c.Open,
			High:   c.High,
			Low:    c.Low,
			Close:  c.Close,
			Volume: int64(c.Volume),
		})
	}

	return ohlcs, nil
}
