package models

type Price_Change struct {
	Current_price float64 `json:"current_price"`
	DailyChange   float64 `json:"daily_change"`
}
