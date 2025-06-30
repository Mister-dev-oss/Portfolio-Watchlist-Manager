package services

import (
	"Backend/external"
	"Backend/models"
	"fmt"
	"sync"
	"time"
)

var quoteCache = make(map[string]cachedQuote)
var mu sync.Mutex

type cachedQuote struct {
	data      models.Price_Change
	timestamp time.Time
}

func GetCachedQuote(ticker string) (models.Price_Change, error) {
	mu.Lock()
	defer mu.Unlock()

	if cached, ok := quoteCache[ticker]; ok {
		if time.Since(cached.timestamp) < time.Minute {
			return cached.data, nil
		}
	}

	quote, err := external.FetchLastQuote(ticker)
	if err != nil {
		return models.Price_Change{}, fmt.Errorf("errore fetch quote: %w", err)
	}

	quoteCache[ticker] = cachedQuote{
		data:      quote,
		timestamp: time.Now(),
	}

	return quote, nil
}
