package models

type PortfolioAsset struct {
	Portfolio_id int     `json:"portfolio_id"`
	Asset_id     int     `json:"asset"`
	Units        float64 `json:"units"`
}

type PortfolioAssetTicker struct {
	Portfolio_id int     `json:"portfolio_id"`
	Ticker       string  `json:"ticker"`
	Units        float64 `json:"units"`
}
