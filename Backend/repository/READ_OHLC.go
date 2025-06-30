package repository

import (
	"Backend/models"
	"database/sql"
	"fmt"
	"time"
)

func READ_OHLCbyTicker(db *sql.DB, assetID int) ([]models.OHLC, error) {

	rows, err := db.Query(`
		SELECT timestamp, open, high, low, close, volume 
		FROM ohlc_data 
		WHERE asset_id = ? 
		ORDER BY timestamp
	`, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ohlcData []models.OHLC
	for rows.Next() {
		var ohlc models.OHLC
		var volume float64

		if err := rows.Scan(&ohlc.Date, &ohlc.Open, &ohlc.High, &ohlc.Low, &ohlc.Close, &volume); err != nil {
			return nil, err
		}

		ohlc.Volume = int64(volume)
		ohlcData = append(ohlcData, ohlc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ohlcData, nil
}

func GetLastOHLCDate(db *sql.DB, assetID int) (*time.Time, error) {
	var dateStr sql.NullString
	err := db.QueryRow(`
		SELECT MAX(timestamp) FROM ohlc_data WHERE asset_id = ?
	`, assetID).Scan(&dateStr)
	if err != nil {
		return nil, err
	}
	if !dateStr.Valid {
		return nil, nil
	}

	parsedDate, err := time.Parse("2006-01-02", dateStr.String)
	if err != nil {
		return nil, fmt.Errorf("errore parsing data: %w", err)
	}

	return &parsedDate, nil
}

func READ_OHLCbyTickerFromDate(db *sql.DB, assetID int, startDate time.Time) ([]models.OHLC, error) {

	rows, err := db.Query(`
		SELECT timestamp, open, high, low, close, volume 
		FROM ohlc_data 
		WHERE asset_id = ? AND date(timestamp) >= date(?) 
		ORDER BY timestamp
	`, assetID, startDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ohlcData []models.OHLC
	for rows.Next() {
		var ohlc models.OHLC
		var volume float64

		if err := rows.Scan(&ohlc.Date, &ohlc.Open, &ohlc.High, &ohlc.Low, &ohlc.Close, &volume); err != nil {
			return nil, err
		}

		ohlc.Volume = int64(volume)
		ohlcData = append(ohlcData, ohlc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ohlcData, nil
}

func READ_ClosePricesByTicker(db *sql.DB, assetID int) ([]float64, error) {
	rows, err := db.Query(`
		SELECT close 
		FROM ohlc_data 
		WHERE asset_id = ? 
		ORDER BY timestamp
	`, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var closePrices []float64
	for rows.Next() {
		var close float64
		if err := rows.Scan(&close); err != nil {
			return nil, err
		}
		closePrices = append(closePrices, close)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return closePrices, nil
}
