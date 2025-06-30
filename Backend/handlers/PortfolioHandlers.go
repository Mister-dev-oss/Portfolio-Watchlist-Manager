package handlers

import (
	"Backend/data"
	"Backend/models"
	"Backend/repository"
	"Backend/services"
	"Backend/utils"

	"github.com/gofiber/fiber/v2"
)

type ModifyPortfolioAssetRequest struct {
	Ticker        string  `json:"ticker"`
	PortfolioName string  `json:"portfolio_name"`
	Quantity      float64 `json:"quantity"`
}

func AddAssetToPortfolioHandler(c *fiber.Ctx) error {

	req := new(ModifyPortfolioAssetRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERRORE": err.Error()})
	}

	if err := services.ServicesAddAssetToPortfolio(data.DB, req.Ticker, req.PortfolioName, req.Quantity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERRORE": err.Error()})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Asset aggiunto/aggiornato !"})
	}

}

func RemoveAssetFromPortfolioHandler(c *fiber.Ctx) error {
	req := new(ModifyPortfolioAssetRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERRORE": err.Error()})
	}

	if err := services.ServicesRemoveAssetFromPortfolio(data.DB, req.Ticker, req.PortfolioName, req.Quantity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERRORE": err.Error()})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Asset aggiornato!"})
	}

}

type GetTicker struct {
	Ticker string `json:"ticker"`
}

func DownloadOhlcForTicker(c *fiber.Ctx) error {

	req := new(GetTicker)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERRORE": err.Error()})
	}

	if err := services.ServicesDownloadOhlcInDB(data.DB, req.Ticker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERRORE": err.Error()})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Ohlc datas downloaddati!"})
	}

}

type PortfolioRequest struct {
	PortfolioName string `json:"portfolio_name"`
}

func RemovePortfolioHandler(c *fiber.Ctx) error {
	portfolio_name := c.Query("portfolio_name")
	if portfolio_name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "Parametro portfolio_name mancante",
		})
	}

	if err := services.ServicesRemovePortfolio(data.DB, portfolio_name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERRORE": err.Error()})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Portfolio rimosso!"})
	}

}

func CreatePortfolioHandler(c *fiber.Ctx) error {
	req := new(PortfolioRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "JSON non valido: " + err.Error(),
		})
	}

	// Validazione del campo obbligatorio
	if req.PortfolioName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "portfolio_name è obbligatorio",
		})
	}

	if err := services.ServicesCreatePortfolio(data.DB, req.PortfolioName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Portfolio creato!",
	})
}

func GetLastQuote(c *fiber.Ctx) error {

	ticker := c.Query("ticker")
	if ticker == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "Parametro 'ticker' mancante",
		})
	}

	quote, err := services.GetCachedQuote(ticker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(quote)
}

func ReadOhlcwithIndicators(c *fiber.Ctx) error {
	ticker := c.Query("ticker")
	if ticker == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "Parametro 'ticker' mancante",
		})
	}

	err := services.ServicesDownloadOhlcInDB(data.DB, ticker)
	if err != nil {
		if err.Error() != "errore HTTP: 403 Forbidden" && err.Error() != "errore HTTP: 429 Too Many Requests" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ERRORE": err.Error(),
			})
		}
	}

	quote, err := services.FetchOHLCbyTicker(data.DB, ticker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	closeQuotes := utils.GetClose(quote)
	indicators := services.GetIndicators(closeQuotes)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"quotes":     quote,
		"indicators": indicators,
	})
}

func GetDispAssets(c *fiber.Ctx) error {

	assetList, err := services.ServicesGetDispAssets(data.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(assetList)
}

func GetAssetInfoFromTicker(c *fiber.Ctx) error {
	ticker := c.Query("ticker")
	if ticker == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "ticker mancante nella query",
		})
	}

	asset, err := services.ServicesGetAssetInfo(data.DB, ticker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(asset)
}

func GetPortfoliosList(c *fiber.Ctx) error {

	portfolios, err := services.ServicesGetPortfolioList(data.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}
	if len(portfolios) == 0 {
		return c.Status(fiber.StatusNoContent).Send(nil)
	}

	return c.Status(fiber.StatusOK).JSON(portfolios)
}

func GetPortfoliosAssets(c *fiber.Ctx) error {
	portfolio_name := c.Query("portfolio_name")
	if portfolio_name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "portfolio mancante nella query",
		})
	}

	id, err := repository.FindPortfolioIdFromName(data.DB, portfolio_name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	assetlist, err := services.GetPortfolioAssetsById(data.DB, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	if len(assetlist) == 0 {
		emptyAssetlist := []models.PortfolioAssetTicker{}
		return c.Status(fiber.StatusOK).JSON(emptyAssetlist)
	}

	return c.Status(fiber.StatusOK).JSON(assetlist)
}

func GetPortfolioAnalysis(c *fiber.Ctx) error {
	portfolio_name := c.Query("portfolio_name")
	if portfolio_name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "portfolio mancante nella query",
		})
	}

	id, err := repository.FindPortfolioIdFromName(data.DB, portfolio_name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	assetlist, err := repository.Read_portfolio_assets(data.DB, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	if len(assetlist) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "Portfolio Vuoto, analisi impossibile",
		})
	}

	var tickers []int
	var quantities []float64
	for _, asset := range assetlist {
		tickers = append(tickers, asset.Asset_id)
		quantities = append(quantities, asset.Units)
	}

	var prices [][]float64
	for _, id := range tickers {
		ticker, err := repository.FindTickerFromAssetID(data.DB, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ERRORE": err.Error(),
			})

		}

		err = services.ServicesDownloadOhlcInDB(data.DB, ticker)

		if err != nil {
			if err.Error() != "errore HTTP: 403 Forbidden" && err.Error() != "errore HTTP: 429 Too Many Requests" {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"ERRORE": err.Error(),
				})
			}
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"ERRORE": err.Error(),
			})

		}

		if data, err := repository.READ_OHLCbyTicker(data.DB, id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ERRORE": err.Error(),
			})
		} else {
			closeData := utils.GetClose(data)
			prices = append(prices, closeData)
		}

	}
	type Metrics struct {
		SharpeRatio float64 `json:"sharpe_ratio"`
		Risk        float64 `json:"risk"`
	}

	sharpe := services.SharpeRatioPortfolio(prices, quantities, 0.0012)
	risk := services.ComputePortfolioRisk(prices, quantities)

	data := Metrics{
		SharpeRatio: sharpe,
		Risk:        risk,
	}

	return c.Status(fiber.StatusOK).JSON(data)

}

func GetAssetsInWatchlist(c *fiber.Ctx) error {

	watchlist, err := services.ServicesGetAssetsFromWatchlist(data.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(watchlist)
}

func AddAssetToWatchList(c *fiber.Ctx) error {
	req := new(GetTicker)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "JSON non valido: " + err.Error(),
		})
	}

	err := services.ServicesAddToWatchlist(data.DB, req.Ticker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func RemoveAssetFromWatchList(c *fiber.Ctx) error {
	ticker := c.Query("ticker")
	if ticker == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "ticker nella query",
		})
	}

	err := services.ServicesRemoveFromWatchlist(data.DB, ticker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetAssetRatings(c *fiber.Ctx) error {
	ticker := c.Query("ticker")
	if ticker == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ERRORE": "ticker nella query",
		})
	}

	data, err := services.FetchOHLCbyTicker(data.DB, ticker)

	ratings, err := services.CalcAssetRatings(data, ticker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ERRORE": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ratings)
}
