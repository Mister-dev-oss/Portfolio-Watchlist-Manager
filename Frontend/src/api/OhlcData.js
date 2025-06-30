import axios from "axios";

export async function GetData(ticker) {
  try {
    const { data } = await axios.get(`http://localhost:3000/api/readohlc?ticker=${ticker}`);

    return data;
    
  } catch (error) {
    console.error("Errore API:", error.response?.data || error.message);
    throw error;
  }
}