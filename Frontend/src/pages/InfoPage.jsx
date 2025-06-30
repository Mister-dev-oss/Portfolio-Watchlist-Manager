import React from 'react';

const InfoPage = () => {
  return (
    <div style={{ maxWidth: 800, margin: "2rem auto", padding: "2rem", background: "#181818", borderRadius: 12, color: "#fff" }}>
      <h1>About This Project</h1>
      <p>
        <strong>Portfolio & Watchlist Manager</strong> is a web application that allows you to:
      </p>
      <ul style={{ marginBottom: "1.5rem" }}>
        <li>Create and manage multiple investment portfolios</li>
        <li>Add, remove, and track assets in your portfolios</li>
        <li>Monitor a personal watchlist of stocks</li>
        <li>View 15-minutes-delayed price changes and basic analytics</li>
        <li>See risk and Sharpe ratio analysis for your portfolios</li>
        <li>Visualize data with interactive charts</li>
      </ul>
      <h2>How to Use</h2>
      <ol style={{ marginBottom: "1.5rem" }}>
        <li>
          <strong>Portfolios:</strong> Go to the Portfolios section to create a new portfolio. Add assets by searching for tickers and specifying quantities. Remove assets or adjust quantities as needed.
        </li>
        <li>
          <strong>Watchlist:</strong> Use the Watchlist to keep an eye on stocks of interest. Add or remove tickers at any time.
        </li>
        <li>
          <strong>Analysis:</strong> For each portfolio, view pie and bar charts of your holdings, and see risk and Sharpe ratio gauges. Click on any asset to view detailed information and price history.
        </li>
      </ol>
      <h2>Disclaimer</h2>
      <p>
        This project was developed as a personal learning exercise to gain experience with full-stack development, including React, Javascript, Go... , and working with financial APIs. It is not intended for production use or for making real investment decisions. Data may be incomplete or delayed, and features may be experimental.
      </p>
      <p style={{ marginTop: "2rem", color: "#aaa" }}>
        Created by Alessandro for educational purposes.
      </p>
    </div>
  );
};

export default InfoPage;
