package repository

import (
	"Backend/models"
	"database/sql"
	"fmt"
)

func InsertOHLCData(db *sql.DB, assetID int, data []models.OHLC) error {

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("errore inizio transazione: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(`
        INSERT INTO ohlc_data (asset_id, timestamp, open, high, low, close, volume)
        VALUES (?, ?, ?, ?, ?, ?, ?)
        ON CONFLICT(asset_id, timestamp) DO NOTHING
    `)
	if err != nil {
		return fmt.Errorf("errore preparazione statement: %w", err)
	}
	defer stmt.Close()

	for _, candle := range data {

		ts := candle.Date

		_, err = stmt.Exec(assetID, ts, candle.Open, candle.High, candle.Low, candle.Close, candle.Volume)
		if err != nil {
			return fmt.Errorf("errore inserimento OHLC: %w", err)
		}
	}
	return nil
}
