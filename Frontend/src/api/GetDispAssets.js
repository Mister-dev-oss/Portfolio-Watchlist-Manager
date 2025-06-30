import axios from "axios";

export async function GetDispAssets() {
  try {
    const { data: tickers } = await axios.get('http://localhost:3000/api/getdispassets');

    if (!Array.isArray(tickers)) {
      console.warn("Dati non in formato atteso:", tickers);
      return [];
    }

    return tickers
  } catch (error) {
    console.error("Errore API:", error.response?.data || error.message);
    return [];
  }
}