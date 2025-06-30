import {
  Chart as ChartJS,
  LineElement,
  LineController,
  CategoryScale,
  TimeScale,
  LinearScale,
  PointElement,
  Tooltip,
  Legend,
  Title
} from 'chart.js';
import { CandlestickController, CandlestickElement } from "chartjs-chart-financial";

import { Chart } from "react-chartjs-2";
import "chartjs-adapter-date-fns";
import { OhlcIndicatorsDataParsing } from "../functions/Data_Parsing";
import { GetData } from "../api/OhlcData";
import { useState, useEffect, useRef } from "react";
import "./TickerGraph.css";
import zoomPlugin from 'chartjs-plugin-zoom';

ChartJS.register(
  LineController,
  CandlestickController,
  LineElement,
  CandlestickElement,
  PointElement,
  CategoryScale,
  TimeScale,
  LinearScale,
  Tooltip,
  Legend,
  Title,
  zoomPlugin
);

export default function OhlcChartWithIndicators({ ticker }) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const chartRef = useRef(null);

  useEffect(() => {
    async function fetchData() {
      try {
        const result = await GetData(ticker);
        setData(result);
        setError(null);
      } catch (err) {
        setError("Errore nel caricamento dei dati");
      }
    }
    fetchData();
  }, [ticker]);

  if (error) return <div className="message">Errore</div>;
  if (!data) return <div className="message">Loading...</div>;

  const { candlestickData = [], indicators = {} } = data ? OhlcIndicatorsDataParsing(data) : {};
  if (candlestickData.length === 0 && Object.keys(indicators).length === 0) {
    return <div className="Message">dati mancanti o corrotti</div>;
  }

  const datasets = [
    {
      label: "OHLC",
      type: "candlestick",
      data: candlestickData,
      borderColor: "white",
      borderWidth: 1,
      color: {
        up: "green",
        down: "red",
        unchanged: "gray",
      },
    },
  ];

  if (indicators.Sma50) {
    datasets.push({
      label: "SMA 50",
      type: "line",
      data: indicators.Sma50.filter(point => point.y > 0),
      borderColor: "blue",
      borderWidth: 1,
      pointRadius: 0,
      hidden: true,
    });
  }

  if (indicators.Sma100) {
    datasets.push({
      label: "SMA 100",
      type: "line",
      data: indicators.Sma100.filter(point => point.y > 0),
      borderColor: "purple",
      borderWidth: 1,
      pointRadius: 0,
      hidden: true,
    });
  }

  if (indicators.Ema20) {
    datasets.push({
      label: "EMA 20",
      type: "line",
      data: indicators.Ema20.filter(point => point.y > 0),
      borderColor: "orange",
      borderWidth: 1,
      pointRadius: 0,
      hidden: true,
    });
  }

  if (indicators.Rsi14) {
    datasets.push({
      label: "RSI 14",
      type: "line",
      data: indicators.Rsi14.filter(point => point.y > 0),
      borderColor: "green",
      borderWidth: 1,
      pointRadius: 0,
      yAxisID: "rsiAxis",
      hidden: true,
    });
  }

  const options = {
    responsive: true,
    maintainAspectRatio: false,
    interaction: {
      mode: 'index',
      intersect: false,
    },
    scales: {
      x: {
        type: "time",
        time: { unit: "day" },
        ticks: { maxRotation: 0 },
      },
      y: {
        title: { display: true, text: "Prezzo" },
        position: "left",
      },
      rsiAxis: {
        position: "right",
        title: { display: true, text: "RSI" },
        min: 0,
        max: 100,
        grid: {
          drawOnChartArea: false,
        },
      },
    },
    plugins: {
      legend: {
        display: true,
        onClick: (e, legendItem, legend) => {
          const ci = legend.chart;
          const index = legendItem.datasetIndex;
          const meta = ci.getDatasetMeta(index);
          meta.hidden = meta.hidden === null ? !ci.data.datasets[index].hidden : null;
          ci.update();
        },
      },
      tooltip: { enabled: true },
      zoom: {
        pan: { enabled: false },
        zoom: {
          wheel: { enabled: false },
          pinch: { enabled: false },
          drag: {
            enabled: true,
            borderColor: "rgba(95, 95, 95, 0.85)",
            borderWidth: 1,
            backgroundColor: "rgba(12, 12, 12, 0.6)",
            modifierKey: null,
          },
          mode: 'x',
          limits: {
            x: { min: 'original', max: 'original' },
            y: { min: 'original', max: 'original' },
          },
        },
      },
    },
  };

  const handleResetZoom = () => {
    if (chartRef.current) {
      chartRef.current.resetZoom();
    }
  };

  return (
    <div style={{ position: "relative", height: "500px" }}>
      <button
        onClick={handleResetZoom}
        style={{
          position: "absolute",
          top: -5,
          right: 10,
          zIndex: 10,
          background: "transparent",
          color:"rgba(129, 129, 129, 0.94)",
          border: "none",
          cursor: "pointer",
          fontSize: "0.9rem",
        }}
      >
        Reset Zoom
      </button>

      <Chart ref={chartRef} type="candlestick" data={{ datasets }} options={options} />
    </div>
  );
}
