package repository

import (
	"Backend/models"
	"database/sql"
)

func GetDisponibleAssets(db *sql.DB) (error, []models.AssetSearch) {
	rows, err := db.Query("SELECT ticker, company_name, exchange,logo FROM assets")
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	// Parsing dei ticker
	var fullData []models.AssetSearch
	for rows.Next() {
		var data models.AssetSearch
		if err := rows.Scan(&data.Ticker, &data.CompanyName, &data.Exchange, &data.Logo); err != nil {
			return err, nil
		}
		fullData = append(fullData, data)
	}
	return nil, fullData
}

func GetAssetInfo(db *sql.DB, ticker string) (error, models.Asset) {
	var asset models.Asset

	err := db.QueryRow("SELECT asset_id, ticker, company_name, industry, description, logo, ceo, exchange, market_cap, sector FROM assets WHERE ticker = ?", ticker).
		Scan(
			&asset.ID,
			&asset.Ticker,
			&asset.CompanyName,
			&asset.Industry,
			&asset.Description,
			&asset.Logo,
			&asset.CEO,
			&asset.Exchange,
			&asset.MarketCap,
			&asset.Sector,
		)
	if err != nil {
		return err, models.Asset{}
	}

	return nil, asset
}
