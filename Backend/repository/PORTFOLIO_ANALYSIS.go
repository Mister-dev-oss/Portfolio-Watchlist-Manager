package repository

import (
	"database/sql"
	"time"
)

func GetPortfolioMinDate(db *sql.DB, portfolioId int) (time.Time, error) {
	// Get all asset IDs from portfolio
	rows, err := db.Query(`
        SELECT asset_id 
        FROM portfolio_assets 
        WHERE portfolio_id = ?`, portfolioId)
	if err != nil {
		return time.Time{}, err
	}
	defer rows.Close()

	var minDate time.Time
	first := true

	// For each asset, get its earliest date
	for rows.Next() {
		var assetId int
		if err := rows.Scan(&assetId); err != nil {
			return time.Time{}, err
		}

		var dateStr string
		err := db.QueryRow(`
            SELECT MIN(timestamp) 
            FROM ohlc_data 
            WHERE asset_id = ?`, assetId).Scan(&dateStr)
		if err != nil {
			return time.Time{}, err
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return time.Time{}, err
		}

		// Update minDate if this is the first asset or if this date is later
		if first || date.After(minDate) {
			minDate = date
			first = false
		}
	}

	return minDate, nil
}
