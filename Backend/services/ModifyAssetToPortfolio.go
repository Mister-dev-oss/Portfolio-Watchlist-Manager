package services

import (
	"Backend/data"
	"Backend/repository"
	"database/sql"
)

func ServicesAddAssetToPortfolio(db *sql.DB, ticker string, portfolioName string, quantity float64) error {

	if err := ServicesDownloadOhlcInDB(db, ticker); err != nil {
		return err
	}

	Pid, Aid, err := repository.FindIdsFromNames(data.DB, ticker, portfolioName)
	if err != nil {
		return err
	}

	if err := repository.AddAssetToPortfolio(data.DB, Pid, Aid, quantity); err != nil {
		return err
	}

	return nil
}

func ServicesRemoveAssetFromPortfolio(db *sql.DB, ticker string, portfolioName string, quantity float64) error {
	Pid, Aid, err := repository.FindIdsFromNames(data.DB, ticker, portfolioName)
	if err != nil {
		return err
	}

	if err := repository.RemoveAssetFromPortfolio(data.DB, Pid, Aid, quantity); err != nil {
		return err
	}

	return nil
}
