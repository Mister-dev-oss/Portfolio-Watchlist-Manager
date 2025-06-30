package main

import (
	data "Backend/data"
	"Backend/external"
	"Backend/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//initialize DB
	data.InitDB()
	defer data.DB.Close()

	//initialize env
	if err := external.Init(); err != nil {
		log.Fatal("Errore nel caricare .env:", err)
	}

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,OPTIONS,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Rotte
	app.Get("/api/getassetratings", handlers.GetAssetRatings)
	app.Get("/api/getportfolioanalysis", handlers.GetPortfolioAnalysis)
	app.Get("/api/getportfolioassets", handlers.GetPortfoliosAssets)
	app.Get("/api/getdispassets", handlers.GetDispAssets)
	app.Get("/api/getportfoliolist", handlers.GetPortfoliosList)
	app.Get("/api/getassetinfo", handlers.GetAssetInfoFromTicker)
	app.Get("/api/readohlc", handlers.ReadOhlcwithIndicators)
	app.Get("/api/getlastquote", handlers.GetLastQuote)
	app.Post("/api/downloadohlc", handlers.DownloadOhlcForTicker)
	app.Post("/api/addasset", handlers.AddAssetToPortfolioHandler)
	app.Post("/api/removeasset", handlers.RemoveAssetFromPortfolioHandler)
	app.Post("/api/createportfolio", handlers.CreatePortfolioHandler)
	app.Delete("/api/removeportfolio", handlers.RemovePortfolioHandler)
	app.Get("/api/getassetinwatchlist", handlers.GetAssetsInWatchlist)
	app.Post("/api/addassetinwatchlist", handlers.AddAssetToWatchList)
	app.Delete("/api/removeassetfromwatchlist", handlers.RemoveAssetFromWatchList)

	if err := app.Listen("127.0.0.1:3000"); err != nil {
		panic(err)
	}
}
