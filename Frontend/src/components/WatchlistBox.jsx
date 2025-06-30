import { useState, useEffect } from 'react';
import { GetWatchlist, AddTickerToWatchlist, RemoveTickerFromWatchlist } from '../api/WatchList';
import './WatchlistBox.css';
import { getQuote } from '../api/TickerPercentage';
import { useNavigate } from 'react-router-dom';
import { MdAdd } from 'react-icons/md';
import AvailableAssetsSearch from './SearchTickerWatchlist';

function WatchlistBox() {
  const [watchlist, setWatchlist] = useState([]);
  const [showInput, setShowInput] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    fetchWatchlist();

    const interval = setInterval(() => {
      fetchWatchlist();
    }, 60000);

    return () => clearInterval(interval);
  }, []);

  async function fetchWatchlist() {
    const data = await GetWatchlist();

    const enriched = await Promise.all(
      data.map(async (asset) => {
        const quote = await getQuote(asset.ticker);

        return {
          ...asset,
          currentPrice: quote?.currentPrice ?? null,
          priceChange: quote?.priceChange ?? null
        };
      })
    );

    setWatchlist(enriched);
  }

  async function handleRemove(ticker) {
    await RemoveTickerFromWatchlist(ticker);
    fetchWatchlist();
  }

  return (
    <div className="watchlist-wrapper">
      <div className="watchlist-header">
        <h2>Your Watchlist</h2>
        <button
          onClick={() => setShowInput(!showInput)}
          className="plus-button"
        >
          {showInput ? (
            <MdAdd size={24} className="rotate-left" />
          ) : (
            <MdAdd size={24} className="rotate-reset" />
          )}
        </button>
      </div>

      <div className={`slide-box ${showInput ? 'visible' : ''}`}>
        {showInput && <AvailableAssetsSearch 
        onClose={() => setShowInput(false)}
        onRefresh={fetchWatchlist}/>}
      </div>

      <div className="watchlist-content">
        <div className="watchlist-box">
          {watchlist.length === 0 ? (
            <p className="NA">No assets</p>
          ) : (
            <ul className="watchlist-ul">
              {watchlist.map((asset, index) => {
                const priceChange = asset.priceChange ?? 0;
                const changeClass =
                  priceChange > 0 ? 'positive' : priceChange < 0 ? 'negative' : 'neutral';

                return (
                  <li
                    key={asset.ticker}
                    style={{ animationDelay: `${index * 0.4}s` }}
                    className="watchlist-item"
                    onClick={() => navigate(`/ticker/?ticker=${asset.ticker}`)}
                  >
                    <div className="data">
                      <h2 style={{ fontSize: '1.5rem', marginBottom: '0.3rem' }}>{asset.ticker}</h2>
                      <p style={{ fontSize: '0.9rem', color: '#666', margin: 0 }}>
                        {asset.company_name} • {asset.exchange} • {asset.sector} <br />
                        Price: ${asset.currentPrice?.toFixed(2) ?? 'N/A'} •{' '}
                        <span className={changeClass}>
                          {Math.abs(priceChange)}%
                        </span>
                      </p>
                    </div>
                    <button
                      className="bottone"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleRemove(asset.ticker);
                      }}
                    >
                      Remove
                    </button>
                  </li>
                );
              })}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default WatchlistBox;
