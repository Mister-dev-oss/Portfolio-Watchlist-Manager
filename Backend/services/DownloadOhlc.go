package services

import (
	"Backend/external"
	"Backend/repository"
	"database/sql"
	"time"
)

func ServicesDownloadOhlcInDB(db *sql.DB, ticker string) error {
	assetID, err := repository.FindAssetIdFromTicker(db, ticker)
	if err != nil {
		return err
	}

	var fromDate time.Time
	lastDate, err := repository.GetLastOHLCDate(db, assetID)
	if err != nil {
		return err
	}
	if lastDate != nil {
		fromDate = lastDate.AddDate(0, 0, 1)
	} else {
		fromDate = time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	}

	now := time.Now()
	toDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if fromDate.After(toDate) {
		return nil
	}

	data, err := external.FetchOHLC(ticker, fromDate, toDate)
	if err != nil {
		if err.Error() == "errore HTTP: 403 Forbidden" {
			return nil
		}
		return err
	}

	if err := repository.InsertOHLCData(db, assetID, data); err != nil {
		return err
	}

	return nil
}
