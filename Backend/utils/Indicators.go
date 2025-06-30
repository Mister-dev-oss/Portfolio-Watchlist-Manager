package utils

func SMA(prices []float64, period int) []float64 {
	result := make([]float64, len(prices))
	if period <= 0 || len(prices) < period {
		return result
	}

	var sum float64
	for i := 0; i < len(prices); i++ {
		sum += prices[i]
		if i >= period {
			sum -= prices[i-period]
		}
		if i >= period-1 {
			result[i] = sum / float64(period)
		} else {
			result[i] = 0
		}
	}
	return result
}

func EMA(prices []float64, period int) []float64 {
	result := make([]float64, len(prices))
	if period <= 0 || len(prices) < period {
		return result
	}

	smoothing := 2.0 / (float64(period) + 1.0)

	var sum float64
	for i := 0; i < period; i++ {
		sum += prices[i]
		result[i] = 0
	}
	emaPrev := sum / float64(period)
	result[period-1] = emaPrev

	for i := period; i < len(prices); i++ {
		ema := (prices[i]-emaPrev)*smoothing + emaPrev
		result[i] = ema
		emaPrev = ema
	}

	return result
}

func RSI(prices []float64, period int) []float64 {
	if len(prices) < period+1 {
		return make([]float64, len(prices))
	}

	rsi := make([]float64, len(prices))
	for i := range rsi {
		rsi[i] = 0
	}

	var gainSum, lossSum float64
	for i := 1; i <= period; i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gainSum += change
		} else {
			lossSum -= change
		}
	}

	avgGain := gainSum / float64(period)
	avgLoss := lossSum / float64(period)

	if avgLoss == 0 {
		for i := period; i < len(prices); i++ {
			rsi[i] = 100
		}
		return rsi
	}

	rs := avgGain / avgLoss
	rsi[period] = 100 - (100 / (1 + rs))

	for i := period + 1; i < len(prices); i++ {
		change := prices[i] - prices[i-1]

		if change > 0 {
			avgGain = (avgGain*float64(period-1) + change) / float64(period)
			avgLoss = (avgLoss * float64(period-1)) / float64(period)
		} else {
			avgGain = (avgGain * float64(period-1)) / float64(period)
			avgLoss = (avgLoss*float64(period-1) - change) / float64(period)
		}

		if avgLoss == 0 {
			rsi[i] = 100
		} else {
			rs = avgGain / avgLoss
			rsi[i] = 100 - (100 / (1 + rs))
		}
	}

	return rsi
}
