package utils

import (
	"math"
)

func PortfolioRisk(weights []float64, returns [][]float64) float64 {
	covMatrix := CovarianceMatrix(returns)
	n := len(weights)
	if len(covMatrix) != n || len(covMatrix[0]) != n {
		return 0
	}

	variance := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			variance += weights[i] * covMatrix[i][j] * weights[j]
		}
	}

	return math.Sqrt(variance)
}
