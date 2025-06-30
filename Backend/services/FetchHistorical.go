package services

import (
	"Backend/models"
	"Backend/repository"
	"database/sql"
)

func FetchOHLCbyTicker(db *sql.DB, ticker string) ([]models.OHLC, error) {
	tickerId, err := repository.FindAssetIdFromTicker(db, ticker)
	if err != nil {
		return []models.OHLC{}, err
	}
	data, err := repository.READ_OHLCbyTicker(db, tickerId)
	return data, err
}
