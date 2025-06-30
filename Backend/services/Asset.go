package services

import (
	"Backend/models"
	"Backend/repository"
	"database/sql"
)

func ServicesGetDispAssets(db *sql.DB) ([]models.AssetSearch, error) {
	if err, list := repository.GetDisponibleAssets(db); err != nil {
		return nil, err
	} else {
		return list, nil
	}

}

func ServicesGetAssetInfo(db *sql.DB, ticker string) (models.Asset, error) {
	if err, asset := repository.GetAssetInfo(db, ticker); err != nil {
		return models.Asset{}, err
	} else {
		return asset, nil
	}
}
