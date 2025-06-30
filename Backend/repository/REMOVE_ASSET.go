package repository

import (
	"Backend/models"
	"database/sql"
	"fmt"
	"math"
)

func RemoveAssetFromPortfolio(db *sql.DB, portfolioId int, assetID int, quantity float64) error {

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

		return fmt.Errorf("L'asset non è stato trovato")

	case nil:
		// Esiste, aggiorniamo
		if quantity < 0 {
			quantity = -quantity
		}
		newQuantity := existing.Units - quantity

		if newQuantity < 0 {
			return fmt.Errorf("impossibile inserire valori maggiori dei precedenti")
		}

		epsilon := 0.001
		if math.Abs(newQuantity) < epsilon {
			_, err := tx.Exec(`
                DELETE FROM portfolio_assets
                WHERE portfolio_id = ? AND asset_id = ?
                `, portfolioId, assetID)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("errore durante la cancellazione dell'asset: %w", err)
			}
			if err := tx.Commit(); err != nil {
				return fmt.Errorf("errore nel commit: %w", err)
			}
			return nil

		}

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
