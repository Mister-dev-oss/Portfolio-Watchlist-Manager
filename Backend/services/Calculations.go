package services

func CalculateUnits(invested float64, price float64) float64 {
	if price == 0 {
		return 0
	}
	return invested / price
}
