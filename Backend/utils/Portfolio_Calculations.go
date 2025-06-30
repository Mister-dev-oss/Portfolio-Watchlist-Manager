package utils

import (
	"math"
)

func Mean(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func StdDev(returns []float64) float64 {
	if len(returns) == 0 {
		return 0
	}
	mean := Mean(returns)
	var variance float64
	for _, r := range returns {
		variance += (r - mean) * (r - mean)
	}
	variance /= float64(len(returns))
	return math.Sqrt(variance)
}

func Covariance(x, y []float64) float64 {
	if len(x) != len(y) || len(x) < 2 {
		return 0
	}
	meanX := Mean(x)
	meanY := Mean(y)
	cov := 0.0
	for i := range x {
		cov += (x[i] - meanX) * (y[i] - meanY)
	}
	return cov / float64(len(x)-1)
}

func CovarianceMatrix(returns [][]float64) [][]float64 {
	n := len(returns)
	covMatrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		covMatrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			covMatrix[i][j] = Covariance(returns[i], returns[j])
		}
	}
	return covMatrix
}
