package models

type Asset struct {
	ID          int     `json:"id"`
	Ticker      string  `json:"ticker"`
	CompanyName string  `json:"company_name"`
	Industry    string  `json:"industry"`
	Description string  `json:"description"`
	Logo        string  `json:"logo"`
	CEO         string  `json:"ceo"`
	Exchange    string  `json:"exchange"`
	MarketCap   float64 `json:"market_cap"`
	Sector      string  `json:"sector"`
}

type AssetSearch struct {
	Ticker      string `json:"ticker"`
	CompanyName string `json:"company_name"`
	Exchange    string `json:"exchange"`
	Logo        string `json:"logo"`
}

type AssetRating struct {
	Ticker     string  `json:"ticker"`
	Volatility float64 `json:"volatility"`
	ATR        float64 `json:"atr"`
}
