package utils

import "math"

func CalculateAssetReturns(prices []float64) []float64 {
	if len(prices) < 2 {
		return nil
	}
	returns := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		if prices[i-1] == 0 {
			returns[i-1] = 0
		} else {
			returns[i-1] = (prices[i] - prices[i-1]) / prices[i-1]
		}
	}
	return returns
}

func AssetLogReturns(prices []float64) []float64 {
	if len(prices) < 2 {
		return nil
	}
	returns := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		if prices[i-1] == 0 {
			returns[i-1] = 0
		} else {
			returns[i-1] = math.Log(prices[i] / prices[i-1])
		}
	}
	return returns
}

func CalculateWeights(quantities []float64, prices []float64) []float64 {
	if len(quantities) != len(prices) || len(quantities) == 0 {
		return nil
	}
	totalValue := 0.0
	for i := range quantities {
		totalValue += quantities[i] * prices[i]
	}
	if totalValue == 0 {
		return nil
	}
	weights := make([]float64, len(quantities))
	for i := range quantities {
		weights[i] = (quantities[i] * prices[i]) / totalValue
	}
	return weights
}

func AggregatePortfolioReturns(assetReturns [][]float64, weights []float64) []float64 {
	numDays := len(assetReturns[0])
	portfolioReturns := make([]float64, numDays)
	for i := 0; i < numDays; i++ {
		for j := 0; j < len(weights); j++ {
			portfolioReturns[i] += assetReturns[j][i] * weights[j]
		}
	}
	return portfolioReturns
}
