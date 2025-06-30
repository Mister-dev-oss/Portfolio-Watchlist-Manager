package utils

import (
	"Backend/models"
)

func GetClose(Data []models.OHLC) []float64 {
	closes := make([]float64, len(Data))
	for i, c := range Data {
		closes[i] = c.Close
	}
	return closes
}
