import { useState, useEffect } from "react";
import {
  GetAssetsInPortfolio,
  RemoveAssetsInPortfolio,
} from "../api/PortfolioAssets";
import "./AssetBox.css";
import { useNavigate, useSearchParams } from "react-router-dom";
import { MdAdd } from "react-icons/md";
import AvailableAssetsSearchPortfolio from "./SearchTickerPortfolio";

function AssetBox({ fatherRefresh }) {
  const [showAddAssetInput, setShowAddAssetInput] = useState(false);
  const navigate = useNavigate();
  const [assets, setAssets] = useState([]);
  const [searchParams] = useSearchParams();
  const portfolioName = searchParams.get("portfolio_name") || "defaultname";
  const [error, setError] = useState(null);
  const [removeQuantities, setRemoveQuantities] = useState(10);
  const [addquantity, setAddQuantity] = useState(0);

  useEffect(() => {
    fetchAssets();
  }, [portfolioName]);

  async function fetchAssets() {
    try {
      const portfolioAssets = await GetAssetsInPortfolio(portfolioName);
      setAssets(portfolioAssets);
    } catch (err) {
      setError(err);
    }
  }

  async function handleRemove(ticker) {
    const qty = removeQuantities[ticker] || 0;
    await RemoveAssetsInPortfolio(portfolioName, ticker, qty);
    setRemoveQuantities((prev) => ({ ...prev, [ticker]: 0 }));
    fetchAssets();
    fatherRefresh();
  }

  return (
    <div className="assetbox-wrapper">
      <div className="assetbox-header">
        <div className="titolation">
          <h2>Portfolio: {portfolioName}</h2>
          <p>Your Assets:</p>
        </div>
        <div className="assetbox-subheader">
          <input
            type="number"
            min="0"
            step="0.01"
            value={addquantity}
            onChange={(e) => {
              const value = e.target.value;
              if (value === "") {
                setAddQuantity("");
              } else {
                const floatValue = parseFloat(value);
                if (!isNaN(floatValue) && floatValue >= 0) {
                  setAddQuantity(floatValue);
                }
              }
            }}
            onKeyDown={(e) => {
              if (["e", "E", "+", "-", ","].includes(e.key)) {
                e.preventDefault();
              }
            }}
            className="quantity-input"
            placeholder="Qty"
          />
          <button
            onClick={() => setShowAddAssetInput(!showAddAssetInput)}
            className="plus-button2"
          >
            {showAddAssetInput ? (
              <MdAdd size={24} className="rotate-left" />
            ) : (
              <MdAdd size={24} className="rotate-reset" />
            )}
          </button>
        </div>
      </div>

      <div className={`assetbox-slide ${showAddAssetInput ? "visible" : ""}`}>
        {showAddAssetInput && (
          <AvailableAssetsSearchPortfolio
            quantity={addquantity}
            portfolio_name={portfolioName}
            onClose={() => setShowAddAssetInput(false)}
            onRefresh={() => {
              fetchAssets();
              fatherRefresh();
              setAddQuantity(0);
            }}
          />
        )}
      </div>

      <div className="assetbox-content">
        <div className="assetbox-list">
          {assets.length === 0 ? (
            <p className="no-assets">No assets found</p>
          ) : (
            <ul className="assetbox-ul">
              {assets.map((asset, index) => (
                <li
                  key={asset.ticker}
                  style={{ animationDelay: `${index * 0.4}s` }}
                  className="assetbox-item"
                  onClick={() => navigate(`/ticker/?ticker=${asset.ticker}`)}
                >
                  <div className="asset-data">
                    <h2>{asset.ticker}</h2>
                    <p>Units held: {asset.units}</p>
                  </div>
                  <div
                    style={{
                      display: "flex",
                      flexDirection: "row",
                      alignItems: "center",
                      gap: "8px",
                      marginLeft: "auto",
                    }}
                  >
                    <input
                      type="number"
                      min="0"
                      step="0.01"
                      value={removeQuantities[asset.ticker] ?? ""}
                      onClick={(e) => e.stopPropagation()}
                      onChange={(e) => {
                        const value = e.target.value;
                        const floatValue = parseFloat(value);
                        setRemoveQuantities((prev) => ({
                          ...prev,
                          [asset.ticker]:
                            value === ""
                              ? ""
                              : !isNaN(floatValue) && floatValue >= 0
                              ? floatValue
                              : 0,
                        }));
                      }}
                      onKeyDown={(e) => {
                        if (["e", "E", "+", "-", ","].includes(e.key)) {
                          e.preventDefault();
                        }
                      }}
                      className="quantity-input"
                      style={{ width: "60px", fontSize: "0.8rem" }}
                      placeholder="Qty"
                    />
                    <button
                      className="bottone"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleRemove(asset.ticker);
                      }}
                    >
                      Remove
                    </button>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default AssetBox;
