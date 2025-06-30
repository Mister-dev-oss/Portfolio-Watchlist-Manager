import { React, useEffect, useState } from "react";
import { GetAssetsInPortfolio } from "../api/PortfolioAssets";
import { useSearchParams } from "react-router-dom";
import AssetBox from "../components/AssetBox";
import PortfolioBox from "../components/PortfolioBox";
import {
  DynamicPieChart,
  DynamicBarChart,
  generateColors,
} from "../components/ChartAnalsisPortfolio";
import SemiCircularGauge from "../components/ChartGauge";
import { GetPortfolioAnalysis } from "../api/Portfolio";
import "./PortfolioPage.css";

const PortfolioPage = () => {
  const [searchparams] = useSearchParams();
  const name = searchparams.get("portfolio_name");
  const [data, setData] = useState([]);
  const [analysis, setAnalysis] = useState(null);
  const [analysisError, setAnalysisError] = useState(null);

  useEffect(() => {
    setAnalysis(null);
    setAnalysisError(null);
    fetchAssets();
    fetchAnalysis();
  }, [name]);

  function handleRefresh() {
    fetchAssets();
    fetchAnalysis();
  }

  async function fetchAssets() {
    try {
      const assets = await GetAssetsInPortfolio(name);
      console.log("Fetched assets", assets);
      setData(assets);
    } catch (err) {
      console.error("Error fetching assets", err);
    }
  }

  async function fetchAnalysis() {
    try {
      const values = await GetPortfolioAnalysis(name);
      console.log("Fetched analysis", values);
      setAnalysis(values);
      setAnalysisError(null);
    } catch (err) {
      console.error("Error fetching analysis", err);
      setAnalysisError(err);
    }
  }

  const names = data.map((item) => item.ticker);
  const quantities = data.map((item) => item.units);
  const colors = generateColors(names.length);

  if (!name) {
    return (
      <div>
        <div style={{fontSize:30, marginBottom:"2rem", color:"white", fontWeight:"bold"}}>Choose a portfolio</div>
        <PortfolioBox />
      </div>
    );
  }

  if (data.length === 0) {
    return (
      <div>
        <p>Add assets to view analysis</p>
        <AssetBox fatherRefresh={handleRefresh} />
      </div>
    );
  }

  return (
    <div>
      <div className="PortfolioPage-container">
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            width: "100%",
          }}
        >
          <AssetBox fatherRefresh={handleRefresh} />
        </div>
        <div className="PortfolioPage-container-item">
          <DynamicPieChart labels={names} data={quantities} colors={colors} />
        </div>
      </div>

      <div
        className="PortfolioPage-container2"
        style={{ marginBottom: "2rem" }}
      >
        <div className="PortfolioPage-container-item2">
          <DynamicBarChart
            labels={names}
            data={quantities}
            colors={colors}
            labelname={"Risk Contribution"}
          />
        </div>

        <div
          style={{
            display: "flex",
            justifyContent: "center",
            gap: "40px",
            alignItems: "center",
          }}
        >
          {analysisError ? (
            analysisError.message.includes("429") ? (
              <p style={{ color: "#70707094" }}>
                Due to an API pricing bottleneck, you'll need to wait 1 minute to get risk scoring.
                If you have more than 5 tickers in your portfolio, you will not get one.
              </p>
            ) : (
              <p style={{ color: "#70707094" }}>
                Error fetching analysis: {analysisError.message}
              </p>
            )
          ) : analysis ? (
            <>
              <SemiCircularGauge
                value={Math.round(analysis.risk * 100) / 100}
                min={0.0}
                max={50.0}
                centerText="risk"
              />
              <SemiCircularGauge
                value={Math.round(analysis.sharpe_ratio * 100) / 100}
                min={0.0}
                max={3.0}
                centerText="sharpe"
              />
            </>
          ) : (
            <p>Loading analysis...</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default PortfolioPage;
