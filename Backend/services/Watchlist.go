package services

import (
	"Backend/repository"
	"database/sql"
)

func ServicesAddToWatchlist(db *sql.DB, ticker string) error {
	return repository.AddAssetToWatchList(db, ticker)
}

func ServicesRemoveFromWatchlist(db *sql.DB, ticker string) error {
	return repository.RemoveAssetFromWatchList(db, ticker)
}

func ServicesGetAssetsFromWatchlist(db *sql.DB) ([]string, error) {
	return repository.GetAssetsFromWatchlist(db)
}
