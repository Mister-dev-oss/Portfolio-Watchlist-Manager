import React from "react";
import {
  Chart as ChartJS,
  ArcElement,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend,
} from "chart.js";
import { Pie, Bar } from "react-chartjs-2";

ChartJS.register(
  ArcElement,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend
);

export function DynamicPieChart({ labels, data, colors }) {
  const pieData = {
    labels,
    datasets: [
      {
        data,
        backgroundColor: colors,
        borderColor: "#fff",
        borderWidth: 1,
      },
    ],
  };

  const barOptions = {
    responsive: true,
    plugins: {
      legend: {
        display: true,
        position: "bottom" , 
        labels: {
        padding: 20,  // spazio tra ogni elemento legenda
      },
      },
  }
}

  return (
    <div style={{ width: "500px" }}>
      <Pie data={pieData} options={barOptions} />
    </div>
  );
}

export function DynamicBarChart({ labels, data, colors, labelname }) {
  const total = data.reduce((acc, curr) => acc + curr, 0);
  const newdata = data.map(
    (item) => Math.round((item / total) * 100 * 100) / 100
  );

  const barData = {
    labels,
    datasets: [
      {
        label: labelname,
        data: newdata,
        backgroundColor: colors,
        borderColor: "#fff",
        borderWidth: 1,
        barPercentage: 0.6,
        categoryPercentage: 0.8,
      },
    ],
  };

  const barOptions = {
    responsive: true,
    plugins: {
      legend: {
        display: false, 
      },
      title: {
        display: true,
        text: labelname,
        position: "top", 
        align: "center", 
        font: {
          size: 20,
        },
        padding: {
          top: 10,
          bottom: 30,
        },
      },
      tooltip: { enabled: true },
    },
    scales: {
      x: { grid: { display: false } },
      y: { beginAtZero: true },
    },
  };

  return (
    <div style={{ width: "600px", margin: "50px auto" }}>
      <Bar data={barData} options={barOptions} />
    </div>
  );
}

export function generateColors(length) {
  return Array.from(
    { length },
    () => `hsl(${Math.floor(Math.random() * 360)}, 70%, 60%)`
  );
}
