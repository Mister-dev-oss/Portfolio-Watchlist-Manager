package services

import (
	"Backend/models"
	"Backend/repository"
	"Backend/utils"
	"database/sql"
	"math"
)

func ComputePortfolioRisk(prices [][]float64, quantities []float64) float64 {
	latestPrices := make([]float64, len(prices))
	returns := [][]float64{}

	for i, assetPrices := range prices {
		latestPrices[i] = assetPrices[len(assetPrices)-1]
		returns = append(returns, utils.CalculateAssetReturns(assetPrices))
	}

	weights := utils.CalculateWeights(quantities, latestPrices)
	risk := utils.PortfolioRisk(weights, returns)
	return risk * 100
}

func SharpeRatioPortfolio(prices [][]float64, quantities []float64, annualRiskFreeRate float64) float64 {
	allReturns := [][]float64{}
	latestPrices := make([]float64, len(prices))

	for i, assetPrices := range prices {
		allReturns = append(allReturns, utils.CalculateAssetReturns(assetPrices))
		latestPrices[i] = assetPrices[len(assetPrices)-1]
	}

	weights := utils.CalculateWeights(quantities, latestPrices)
	portfolioReturns := utils.AggregatePortfolioReturns(allReturns, weights)

	meanReturn := utils.Mean(portfolioReturns)
	std := utils.StdDev(portfolioReturns)

	if std == 0 {
		return 0
	}

	dailyRf := annualRiskFreeRate / 252
	dailySharpe := (meanReturn - dailyRf) / std

	return dailySharpe * math.Sqrt(252)
}

func GetIndicators(prices []float64) models.Indicators {
	var Indicators models.Indicators
	Indicators.Sma50 = utils.SMA(prices, 50)
	Indicators.Sma100 = utils.SMA(prices, 100)
	Indicators.Rsi14 = utils.RSI(prices, 14)
	Indicators.Ema20 = utils.EMA(prices, 20)
	return Indicators

}

func GetPortfolioAssetsById(db *sql.DB, id int) ([]models.PortfolioAssetTicker, error) {
	if assets, err := repository.Read_portfolio_assets_tickerModel(db, id); err != nil {
		return []models.PortfolioAssetTicker{}, err
	} else {
		return assets, nil
	}
}

func CalcAssetRatings(data []models.OHLC, ticker string) (models.AssetRating, error) {
	closes := utils.GetClose(data)
	Histvol := utils.HistoricalVolatility(closes)
	Atr := utils.ATRFullPeriod(data)
	Rating := models.AssetRating{Ticker: ticker, Volatility: Histvol, ATR: Atr}
	return Rating, nil
}
