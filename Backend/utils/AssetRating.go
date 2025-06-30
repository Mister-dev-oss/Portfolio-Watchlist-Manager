package utils

import (
	"Backend/models"
	"math"
)

func HistoricalVolatility(prices []float64) float64 {
	returns := AssetLogReturns(prices)
	if len(returns) == 0 {
		return 0
	}
	std := StdDev(returns)
	return std * math.Sqrt(252)
}

func ATRFullPeriod(ohlc []models.OHLC) float64 {
	n := len(ohlc)
	if n < 2 {
		return 0
	}

	tr := make([]float64, n-1)

	for i := 1; i < n; i++ {
		h := ohlc[i].High
		l := ohlc[i].Low
		cPrev := ohlc[i-1].Close

		range1 := h - l
		range2 := math.Abs(h - cPrev)
		range3 := math.Abs(l - cPrev)

		tr[i-1] = math.Max(range1, math.Max(range2, range3))
	}

	var sum float64
	for _, val := range tr {
		sum += val
	}

	atr := sum / float64(len(tr))
	return atr
}
