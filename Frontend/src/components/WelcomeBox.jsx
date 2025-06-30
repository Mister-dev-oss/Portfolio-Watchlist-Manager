import React, { useState, useEffect } from "react";
import "./WelcomeBox.css"
import { getQuote } from '../api/TickerPercentage';


function WelcomeBox() {
  const tickers = ["AAPL", "NVDA", "TSLA"];
  const [quotes, setQuotes] = useState({});

  useEffect(() => {
    let isMounted = true;

    async function fetchQuotes() {
      const newQuotes = {};
      for (const ticker of tickers) {
        const data = await getQuote(ticker);
        console.log(ticker, data);
        newQuotes[ticker] = data || { currentPrice: "N/A", priceChange: "N/A" };
      }
      if (isMounted) setQuotes(newQuotes);
    }

    fetchQuotes();
    const intervalId = setInterval(fetchQuotes, 60000);

    return () => {
      isMounted = false;
      clearInterval(intervalId);
    };
  }, []);

   return (
    <div className="welcome-box">
      <div className="welcome-box-text">
        <p>Welcome!</p>
        <div className="welcome-box-text2">Manage your portfolios in seconds!</div>
      </div>
      <div className="mini-boxes">
        {tickers.map((ticker) => {
          const priceChange = quotes[ticker]?.priceChange ?? 0;
          const changeClass = priceChange > 0 ? "positive" : priceChange < 0 ? "negative" : "neutral";

          return (
            <div key={ticker} className="mini-box">
              <div className="texts">{ticker}</div>
              <div className="texts">
                <div>{quotes[ticker]?.currentPrice ?? "..."}</div>
                <div className={changeClass}>
                {priceChange ?? "..."}%
                </div>
              </div>
            </div>
          );
        })}
       </div>
    </div>
  );
}

export default WelcomeBox