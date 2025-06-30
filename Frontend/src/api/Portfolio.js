import axios from 'axios';
import { data } from 'react-router-dom';

export async function GetPortfolioList() {
  try {
    const { data } = await axios.get('http://localhost:3000/api/GetPortfolioList');

    if (!Array.isArray(data)) {
      console.warn("Dati non in formato atteso:", data);
      return [];
    }

    return data;
  } catch (error) {
    console.error("Errore API:", error.response?.data || error.message);
    return [];
  }
}

export async function CreatePortfolio(name) {
  try {
    const { data } = await axios.post('http://localhost:3000/api/CreatePortfolio', {
      portfolio_name: name
    });

    console.log("Portfolio Aggiunto:", data.message);
    return true;
  } catch (error) {
    console.error("Errore nella creazione del portfolio:", error.response?.data || error.message);
    return false;
  }
}

export async function RemovePortfolio(name) {
  try {
    const { data } = await axios.delete(`http://localhost:3000/api/RemovePortfolio?portfolio_name=${name}`);

    console.log("Rimozione avvenuta:", data.message);
    return true;
  } catch (error) {
    console.error("Errore nella rimozione:", error.response?.data || error.message);
    return false;
  }
}

export async function GetPortfolioAnalysis(name) {
  
  const { data } = await axios.get(`http://localhost:3000/api/getportfolioanalysis?portfolio_name=${name}`);

  console.log("Analysis obtained", data.message);
  return data;
}


