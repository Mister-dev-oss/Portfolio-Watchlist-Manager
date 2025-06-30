package repository

import (
	"Backend/models"
	"database/sql"
	"fmt"
)

func AddAssetToPortfolio(db *sql.DB, portfolioId int, assetID int, quantity float64) error {

	if quantity <= 0 {
		return fmt.Errorf("impossibile inserire valori negativi")
	}

	var existing models.PortfolioAsset
	err := db.QueryRow(`
		SELECT quantity
		FROM portfolio_assets 
		WHERE asset_id = ? AND portfolio_id = ?
	`, assetID, portfolioId).Scan(
		&existing.Units,
	)

	tx, db_err := db.Begin()
	if db_err != nil {
		return fmt.Errorf("errore durante l'inizio della transazione: %w", db_err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	switch err {
	case sql.ErrNoRows:
		// Primo inserimento per quell’asset in quel portafoglio
		_, err = tx.Exec(`
			INSERT INTO portfolio_assets (portfolio_id, asset_id, quantity)
			VALUES (?, ?, ?)
		`, portfolioId, assetID, quantity)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("errore nell'inserimento dell'asset: %w", err)
		}
	case nil:
		// Esiste già, aggiorniamo
		newQuantity := existing.Units + quantity

		_, err = tx.Exec(`
			UPDATE portfolio_assets
			SET quantity = ?
			WHERE portfolio_id = ? AND asset_id = ?
		`, newQuantity, portfolioId, assetID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("errore nell'aggiornamento dell'asset: %w", err)
		}
	default:
		tx.Rollback()
		return fmt.Errorf("errore nella lettura dell’asset nel portafoglio: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("errore nel commit: %w", err)
	}
	return nil
}
