package repository

import (
	"database/sql"
	"fmt"
)

func FindAssetIdFromTicker(db *sql.DB, ticker string) (int, error) {

	var assetID int

	err := db.QueryRow("SELECT asset_id FROM assets WHERE ticker = ?", ticker).Scan(&assetID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("ticker %s non trovato nel database", ticker)
		}
		return 0, fmt.Errorf("errore nella ricerca dell'asset_id: %w", err)
	}
	return assetID, nil
}

func FindTickerFromAssetID(db *sql.DB, assetID int) (string, error) {

	var ticker string

	err := db.QueryRow("SELECT ticker FROM assets WHERE asset_id = ?", assetID).Scan(&ticker)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("ticker con asset_id %d non trovato nel database", assetID)
		}
		return "", fmt.Errorf("errore nella ricerca dell'asset_id: %w", err)
	}
	return ticker, nil
}

func FindPortfolioIdFromName(db *sql.DB, name string) (int, error) {
	var portfolioId int
	err := db.QueryRow("SELECT id FROM portfolios WHERE name = ?", name).Scan(&portfolioId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("portfolioName %s non trovato nel database", name)
		}
		return 0, fmt.Errorf("errore nella ricerca del portfolioId: %w", err)
	}
	return portfolioId, nil
}

func FindIdsFromNames(db *sql.DB, assetTicker string, portfolioName string) (int, int, error) {
	var assetID int

	err := db.QueryRow("SELECT asset_id FROM assets WHERE ticker = ?", assetTicker).Scan(&assetID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, fmt.Errorf("ticker %s non trovato nel database", assetTicker)
		}
		return 0, 0, fmt.Errorf("errore nella ricerca dell'asset_id: %w", err)
	}

	var portfolioId int
	err = db.QueryRow("SELECT id FROM portfolios WHERE name = ?", portfolioName).Scan(&portfolioId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, fmt.Errorf("portfolioName %s non trovato nel database", portfolioName)
		}
		return 0, 0, fmt.Errorf("errore nella ricerca del portfolioId: %w", err)
	}

	return portfolioId, assetID, nil
}
