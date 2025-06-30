package services

import (
	"Backend/repository"
	"database/sql"
)

func ServicesCreatePortfolio(db *sql.DB, name string) error {
	if err := repository.CreatePortfolio(db, name); err != nil {
		return err
	}
	return nil
}

func ServicesRemovePortfolio(db *sql.DB, name string) error {

	Pid, err := repository.FindPortfolioIdFromName(db, name)
	if err != nil {
		return err
	}
	if err := repository.RemovePortfolio(db, Pid); err != nil {
		return err
	}
	return nil
}

func ServicesGetPortfolioList(db *sql.DB) ([]string, error) {
	if err, list := repository.GetPortfolioList(db); err != nil {
		return []string{}, err
	} else {
		return list, nil
	}
}
