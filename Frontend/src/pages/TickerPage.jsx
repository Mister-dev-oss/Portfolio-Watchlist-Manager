import { React, useEffect, useState } from "react";
import OhlcChartWithIndicators from "../components/TickerGraph";
import { useSearchParams } from "react-router-dom";
import "./TickerPage.css";
import { GetAssetInfo, GetAssetRatings } from "../api/WatchList";
import { getQuote } from "../api/TickerPercentage";

const TickerPage = () => {
  const [searchparams] = useSearchParams();
  const ticker = searchparams.get("ticker") || "defaultTicker";
  const [AssetInfo, setAssetInfo] = useState(null);
  const [AssetRatings, setAssetRatings] = useState(null);
  const [Value, setValue] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    let isMounted = true;

    async function fetchRatingsWithRetry(ticker, retries = 3, delay = 1000) {
      for (let attempt = 0; attempt < retries; attempt++) {
        const data = await GetAssetRatings(ticker);
        if (data.atr !== 0 || data.volatility !== 0) {
          return data;
        }
        await new Promise((res) => setTimeout(res, delay));
      }
      return await GetAssetRatings(ticker);
    }

    async function fetchData() {
      try {
        const [info, data, ratings] = await Promise.all([
          GetAssetInfo(ticker),
          getQuote(ticker),
          fetchRatingsWithRetry(ticker),
        ]);

        if (isMounted) {
          setAssetInfo(info);
          setValue(data);
          let newratings = {
            atr: Number(ratings.atr.toFixed(2)),
            volatility: Number(ratings.volatility.toFixed(2)),
          };
          setAssetRatings(newratings);
          setError(null);
        }
      } catch (err) {
        if (isMounted) {
          setError("Error loading data");
        }
      }
    }

    fetchData();

    return () => {
      isMounted = false;
    };
  }, [ticker]);

  if (error) return <div className="message">Error</div>;
  if (!AssetInfo) return <div className="message">Loading...</div>;

  return (
    <div style={{marginTop:" 2rem"}}>
      <div className="TickerInfo-container">
        <OhlcChartWithIndicators ticker={ticker} />

        <div className="Specifics">
          <div className="Specifics card">
            <div className="specifics-header">
              <h1 className="ticker-symbol">
                {AssetInfo.ticker}
                <span
                  className={`price-change-inline ${
                    Value.priceChange > 0
                      ? "positive"
                      : Value.priceChange < 0
                      ? "negative"
                      : "neutral"
                  }`}
                >
                  {Value.priceChange !== undefined
                    ? ` ${Value.priceChange}%`
                    : " N/A"}
                </span>
              </h1>
              <p className="ticker-name">{AssetInfo.company_name}</p>
            </div>

            <div className="specifics-grid">
              <div className="info-block">
                <h3>Industry</h3>
                <p>{AssetInfo.industry}</p>
              </div>
              <div className="info-block">
                <h3>Exchange</h3>
                <p>{AssetInfo.exchange}</p>
              </div>
              <div className="info-block">
                <h3>Sector</h3>
                <p>{AssetInfo.sector}</p>
              </div>
              <div className="info-block">
                <h3>CEO</h3>
                <p>{AssetInfo.ceo}</p>
              </div>
              <div className="info-block">
                <h3>Market Cap</h3>
                <p>{AssetInfo.market_cap}</p>
              </div>
              <div className="info-block">
                <h3>Current Price</h3>
                <p>{Value.currentPrice || "N/A"}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="BottomBlock">
        <div className="BottomBlock1">
          <div className="info-block">
            <h3>Average True Range</h3>
            <p>{AssetRatings.atr}</p>
          </div>
          <div className="info-block">
            <h3>Historical Volatility</h3>
            <p>{AssetRatings.volatility || "N/A"}</p>
          </div>
        </div>
        <div>
          <div className="info-block">
            <h3>Description</h3>
            <p>{AssetInfo.description}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TickerPage;
