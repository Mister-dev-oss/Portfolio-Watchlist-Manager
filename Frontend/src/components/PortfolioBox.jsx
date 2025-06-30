import { useState, useEffect } from "react";
import {
  GetPortfolioList,
  CreatePortfolio,
  RemovePortfolio,
} from "../api/Portfolio";
import "./PortfolioBox.css";
import { MdClose, MdAdd } from "react-icons/md";
import { useNavigate } from "react-router-dom";

function PortfolioBox() {
  const [portfolios, setPortfolios] = useState([]);
  const [newPortfolioName, setNewPortfolioName] = useState("");
  const [showInput, setShowInput] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    fetchPortfolios();
  }, []);

  async function fetchPortfolios() {
    const data = await GetPortfolioList();
    setPortfolios(data);
  }

  async function handleAdd() {
    if (!newPortfolioName.trim()) return;
    await CreatePortfolio(newPortfolioName);
    setNewPortfolioName("");
    setShowInput(false);
    fetchPortfolios();
  }

  async function handleRemove(name) {
    await RemovePortfolio(name);
    fetchPortfolios();
  }

  return (
    <div className="portfolio-container">
      <div className="portfolio-container-top">
        <h2>Your Portfolios</h2>
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

      {portfolios.length === 0 ? (
        <p className="NA">No portfolios</p>
      ) : (
        <ul className="watchlist-ul">
          {portfolios.map((name, index) => (
            <li
              key={name}
              style={{ animationDelay: `${index * 0.4}s` }}
              onClick={() => navigate(`/portfolio/?portfolio_name=${name}`)}
            >
              <span className="portfolio-name">{name}</span>
              <button
                className="bottone"
                onClick={(e) => {
                  e.stopPropagation();
                  handleRemove(name);
                }}
              >
                Remove
              </button>
            </li>
          ))}
        </ul>
      )}

      {showInput && (
        <div className="add-box">
          <input
            value={newPortfolioName}
            onChange={(e) => setNewPortfolioName(e.target.value)}
            placeholder="New portfolio name"
          />
          <button className="bottone" onClick={handleAdd}>
            Confirm
          </button>
        </div>
      )}
    </div>
  );
}

export default PortfolioBox;
