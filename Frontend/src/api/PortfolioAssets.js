import axios from "axios";

export async function GetAssetsInPortfolio(portfolio_name) {
  try {
    const { data: Assets } = await axios.get(`http://localhost:3000/api/getportfolioassets?portfolio_name=${portfolio_name}`);

    if (!Array.isArray(Assets)) {
      console.warn("Dati non in formato atteso:", Assets);
      return [];
    }

    return Assets
  } catch (error) {
    console.error("Errore API:", error.response?.data || error.message);
    return [];
  }
}

export async function AddAssetsInPortfolio(portfolio_name, ticker, quantity) {
  try {
    const { data } = await axios.post('http://localhost:3000/api/addasset', {
      ticker: ticker,
      portfolio_name: portfolio_name,
      quantity: quantity
    });

    console.log("Asset Aggiunto:", data.message);
    return true;
  } catch (error) {
    console.error("Errore nell'aggiunta dell'asset al portfolio:", error.response?.data || error.message);
    return false;
  }
}

export async function RemoveAssetsInPortfolio(portfolio_name, ticker, quantity) {
    try {
    const { data } = await axios.post('http://localhost:3000/api/removeasset', {
      ticker: ticker,
      portfolio_name: portfolio_name,
      quantity: quantity
    });
  
    console.log("Ticker Rimosso/diminuito:", data.message);
    return true;
  } catch (error) {
    console.error("Errore nella rimozione/diminuzione dell'asset al portfolio:", error.response?.data || error.message);
    return false;
  }
}