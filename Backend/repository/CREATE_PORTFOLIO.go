package repository

import (
	"database/sql"
	"fmt"
	"time"
)

func CreatePortfolio(db *sql.DB, name string) error {
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

	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM portfolios WHERE name = ?)", name).Scan(&exists)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore durante la verifica esistenza portfolio: %w", err)
	}

	if exists {
		tx.Rollback()
		return fmt.Errorf("errore: il portfolio '%s' già esiste", name)
	}

	createdAt := time.Now().UTC()
	_, err = tx.Exec(`
        INSERT INTO portfolios (name, created_at)
        VALUES (?, ?)
    `, name, createdAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("errore inserimento portfolio: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("errore commit transazione: %w", err)
	}

	return nil
}
