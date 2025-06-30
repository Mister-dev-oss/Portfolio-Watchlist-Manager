package repository

import (
	"database/sql"
	"fmt"
)

func AddAssetToWatchList(db *sql.DB, ticker string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("errore inizio transazione: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	var existsInAssets bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM assets WHERE ticker = ?)", ticker).Scan(&existsInAssets)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore durante la verifica esistenza in assets: %w", err)
	}
	if !existsInAssets {
		tx.Rollback()
		return fmt.Errorf("errore: il ticker '%s' non è registrato nel DB", ticker)
	}

	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM watchlist WHERE ticker = ?)", ticker).Scan(&exists)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore durante la verifica esistenza portfolio: %w", err)
	}

	if exists {
		tx.Rollback()
		return fmt.Errorf("errore: il ticker '%s' è già in watchlist", ticker)
	}

	_, err = tx.Exec(`
        INSERT INTO watchlist (ticker)
        VALUES (?)
    `, ticker)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore inserimento ticker in watchlist: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("errore commit transazione: %w", err)
	}

	return nil
}

func RemoveAssetFromWatchList(db *sql.DB, ticker string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("errore inizio transazione: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	res, err := tx.Exec("DELETE FROM watchlist WHERE ticker = ?", ticker)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("errore durante la cancellazione del ticker: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("errore nel recupero delle righe interessate: %w", err)
	}
	if rowsAffected == 0 {
		_ = tx.Rollback()
		return fmt.Errorf("nessun ticker '%s' trovato in watchlist", ticker)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("errore nel commit della transazione: %w", err)
	}

	return nil
}

func GetAssetsFromWatchlist(db *sql.DB) ([]string, error) {

	rows, err := db.Query("SELECT ticker FROM watchlist")
	if err != nil {
		return nil, fmt.Errorf("errore durante la query della watchlist: %w", err)
	}
	defer rows.Close()

	var tickers []string
	for rows.Next() {
		var ticker string
		if err := rows.Scan(&ticker); err != nil {
			return nil, fmt.Errorf("errore nella scansione dei risultati: %w", err)
		}
		tickers = append(tickers, ticker)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("errore durante l'iterazione dei risultati: %w", err)
	}

	return tickers, nil
}
