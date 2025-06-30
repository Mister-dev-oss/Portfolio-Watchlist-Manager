package repository

import (
	"Backend/models"
	"database/sql"
)

func Read_portfolio_assets(db *sql.DB, portfolio_id int) ([]models.PortfolioAsset, error) {

	rows, err := db.Query(`
		SELECT portfolio_id, asset_id, quantity
		FROM portfolio_assets 
		WHERE portfolio_id = ? 
		ORDER BY quantity
	`, portfolio_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var portfolioAssets []models.PortfolioAsset
	for rows.Next() {
		var portfolioAsset models.PortfolioAsset

		if err := rows.Scan(&portfolioAsset.Portfolio_id, &portfolioAsset.Asset_id, &portfolioAsset.Units); err != nil {
			return nil, err
		}
		portfolioAssets = append(portfolioAssets, portfolioAsset)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return portfolioAssets, nil
}

func GetPortfolioList(db *sql.DB) (error, []string) {

	rows, err := db.Query("SELECT name FROM portfolios")
	if err != nil {
		return err, []string{}
	}
	defer rows.Close()

	var portfolios []string
	for rows.Next() {
		var ticker string
		if err := rows.Scan(&ticker); err != nil {
			return err, []string{}
		}
		portfolios = append(portfolios, ticker)
	}
	return nil, portfolios
}

func Read_portfolio_assets_tickerModel(db *sql.DB, portfolio_id int) ([]models.PortfolioAssetTicker, error) {

	rows, err := db.Query(`
		SELECT portfolio_id, asset_id, quantity
		FROM portfolio_assets 
		WHERE portfolio_id = ? 
		ORDER BY quantity
	`, portfolio_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var portfolioAssets []models.PortfolioAssetTicker
	for rows.Next() {
		var portfolioAsset models.PortfolioAsset
		var portfolioAssetTiker models.PortfolioAssetTicker

		if err := rows.Scan(&portfolioAsset.Portfolio_id, &portfolioAsset.Asset_id, &portfolioAsset.Units); err != nil {
			return nil, err
		}
		portfolioAssetTiker.Portfolio_id = portfolioAsset.Portfolio_id
		portfolioAssetTiker.Ticker, err = FindTickerFromAssetID(db, portfolioAsset.Asset_id)
		if err != nil {
			return nil, err
		}
		portfolioAssetTiker.Units = portfolioAsset.Units
		portfolioAssets = append(portfolioAssets, portfolioAssetTiker)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return portfolioAssets, nil
}
