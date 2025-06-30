import React, { useEffect, useState } from 'react';
import { GetDispAssets } from '../api/GetDispAssets';
import { AddTickerToWatchlist } from '../api/WatchList';
import "./SearchTicker.css"

export default function AvailableAssetsSearch({ onClose, onRefresh }) {

 async function handleAdd(ticker) {
    if (!ticker.trim()) return;
    await AddTickerToWatchlist(ticker);
    onRefresh(); 
    onClose();
  } 

  async function loadData() {
    const data = await GetDispAssets();
    setData(data)
    return
  }

  useEffect(() => {
    loadData();
  }, []);


  const [filter, setFilter] = useState('');
  const [data, setData] = useState([])
  const filteredAssets = filter
  ? data.filter(a =>
      a.ticker.toLowerCase().includes(filter.toLowerCase())|| 
      a.company_name.toLowerCase().includes(filter.toLocaleLowerCase())
    )
  : [];

  return (
  <div 
    style={{
      maxHeight: '300px',
      overflowY: 'auto',
      marginTop: '-1.4rem',
      padding: '8px', 
    }}
  >
    <input
      type="text"
      placeholder="Filter by ticker..."
      value={filter}
      onChange={e => setFilter(e.target.value)}
      className="asset-search-input"
    />

    <div className="asset-search-results hide-scrollbar">
      {filteredAssets.length === 0 ? (
        <p className="asset-search-empty">No assets found</p>
      ) : (
        <ul className="asset-search-list"  >
          {filteredAssets.slice(0, 5).map((asset, index) => (
            <li
              onClick={() => handleAdd(asset.ticker)}
              key={asset.ticker}
              className="asset-item"
              style={{ 
                animation: `Fade 0.5s ease-in forwards`,
                animationDelay: `${index * 0.4}s`
              }}
            >
              <div>
                <strong>{asset.ticker}</strong> - {asset.company_name} ({asset.exchange})
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  </div>
);
}