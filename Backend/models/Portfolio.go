package models

import "time"

type Portfolio struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	CreatedAt       time.Time        `json:"created_at"`
	PortfolioAssets []PortfolioAsset `json:"portfolio_assets"`
}

// For a v2 with multiple portfolios
type PortfolioList struct {
	Portfolios []Portfolio `json:"portfolio_list"`
}
