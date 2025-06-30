package repository

import (
	"database/sql"
	"fmt"
)

func RemovePortfolio(db *sql.DB, portfolioId int) error {
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

	res, err := tx.Exec("DELETE FROM portfolios WHERE id = ?", portfolioId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore durante la cancellazione del portfolio: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore ottenendo righe cancellate: %w", err)
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("nessun portfolio trovato con id %d", portfolioId)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("errore nel commit della transazione: %w", err)
	}

	return nil
}
