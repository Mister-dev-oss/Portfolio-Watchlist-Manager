import axios from 'axios';

export async function GetWatchlist() {
  try {
    const { data: tickers } = await axios.get('http://localhost:3000/api/getassetinwatchlist');

    if (!Array.isArray(tickers)) {
      console.warn("Dati non in formato atteso:", tickers);
      return [];
    }

    const assetPromises = tickers.map(async (ticker) => {
      try {
        const response = await axios.get(`http://localhost:3000/api/getassetinfo?ticker=${ticker}`);
        return response.data;
      } catch (error) {
        console.warn(`Errore nel recupero di ${ticker}:`, error.response?.data || error.message);
        return null; 
      }
    });

    
    const assets = await Promise.all(assetPromises);

    
    return assets.filter(asset => asset !== null);
  } catch (error) {
    console.error("Errore API:", error.response?.data || error.message);
    return [];
  }
}


export async function AddTickerToWatchlist(ticker) {
  try {
    const { data } = await axios.post('http://localhost:3000/api/addassetinwatchlist', {
      ticker: ticker
    });

    console.log("Ticker Aggiunto:", data.message);
    return true;
  } catch (error) {
    console.error("Errore nell'aggiunta del ticker alla watchlist:", error.response?.data || error.message);
    return false;
  }
}

export async function RemoveTickerFromWatchlist(ticker) {
  try {
    const { data } = await axios.delete(`http://localhost:3000/api/removeassetfromwatchlist?ticker=${ticker}`);

    console.log("Rimozione avvenuta:", data.message);
    return true;
  } catch (error) {
    console.error("Errore nella rimozione:", error.response?.data || error.message);
    return false;
  }
}


export async function GetAssetInfo(ticker) {
  try {
    const response = await axios.get(`http://localhost:3000/api/getassetinfo?ticker=${ticker}`);
    return response.data;
  } catch (error) {
    console.error("Errore:", error.response?.data || error.message);
    throw error;
  }
}

export async function GetAssetRatings(ticker) {
  try {
    const response = await axios.get(`http://localhost:3000/api/getassetratings?ticker=${ticker}`);
    return response.data;
  } catch (error) {
    console.error("Errore:", error.response?.data || error.message);
    throw error;
  }
}



