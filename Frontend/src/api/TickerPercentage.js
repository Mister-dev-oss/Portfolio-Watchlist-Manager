import axios from 'axios';

export async function getQuote(ticker) {
  try {
    const response = await axios.get(`http://localhost:3000/api/getlastquote?ticker=${ticker}`);
    const { current_price, daily_change } = response.data;

    return {
      currentPrice: Math.round(current_price * 100) / 100,
      priceChange: Math.round(daily_change * 100) / 100
    };
  } catch (error) {
    console.error(`Errore API per ticker ${ticker}:`, error.response?.data || error.message);
    return null;
  }
}


