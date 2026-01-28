# Portfolio & Watchlist Manager

## Overview

**Portfolio & Watchlist Manager** is a full-stack web application for managing investment portfolios and personal stock watchlists. It allows you to create and manage multiple portfolios, add or remove assets, monitor real-time and historical prices, and visualize risk and performance analytics with interactive charts.

This project was developed by Alessandro as a personal learning exercise to gain experience with React, JavaScript, Go, and working with financial APIs.  
**It is not intended for production use or real investment decisions. Data may be incomplete or delayed, and features may be experimental.**

---

## Features

- Create and manage multiple investment portfolios
- Add, remove, and track assets in your portfolios
- Monitor a personal watchlist of stocks
- View 15-minutes-delayed price changes and basic analytics
- See risk and Sharpe ratio analysis for your portfolios
- Visualize data with interactive charts

---

## Project Structure

```
Backend/
  main.go                # Go backend entry point
  handlers/              # HTTP handlers (API endpoints)
  models/                # Data models (Asset, Portfolio, etc.)
  repository/            # Database access logic
  services/              # Business logic (risk, analytics)
  utils/                 # Utility functions
  assets/                # Static assets (logos, etc.)
  data/                  # Database initialization
  external/              # External API integrations (Finnhub, Polygon)
  portfolio.db           # SQLite database
  .env.example           # Example environment file for API keys

Frontend/
  src/
    api/                 # API calls to backend
    components/          # React UI components
    functions/           # Data parsing and helpers
    pages/               # Main pages (Home, Portfolio, Ticker, Info)
    App.jsx              # Main React app
    App.css              # Global styles
  public/                # Static files
  index.html             # HTML entry point
  package.json           # Frontend dependencies and scripts
```

---

## Environment Variables & API Keys

Before running the backend, you **must** set up your environment variables for API access.

1. **Copy the example environment file**  
   In the `Backend` directory, duplicate the `.env.example` file and rename it to `.env`.

2. **Edit your API keys**  
   Open the new `.env` file and replace the placeholders with your actual API keys:
   ```
   FINNHUB_API_KEY=your-finnhub-api-key
   POLYGON_API_KEY=your-polygon-api-key
   ```
---

## How to Use

### Prerequisites

- [Node.js](https://nodejs.org/) (for frontend)
- [Go](https://golang.org/) (for backend)
- Internet connection (for financial data APIs)

### Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/Mister-dev-oss/Portfolio-Watchlist-Manager
   cd Portfolio-Watchlist-Manager
   ```

2. **Backend:**
   - Navigate to the `Backend` folder.
   - **Set up your `.env` file with API keys as described above.**
   - Run the backend server:
     ```sh
     go run main.go
     ```
   - The backend will listen on `127.0.0.1:3000`.

3. **Frontend:**
   - Navigate to the `Frontend` folder.
   - Install dependencies:
     ```sh
     npm install
     ```
   - Start the development server:
     ```sh
     npm run dev
     ```
   - The frontend will be available at [http://localhost:5173](http://localhost:5173).

---

## Usage

- **Portfolios:** Go to the Portfolios section to create a new portfolio. Add assets by searching for tickers and specifying quantities. Remove assets or adjust quantities as needed.
- **Watchlist:** Use the Watchlist to keep an eye on stocks of interest. Add or remove tickers at any time.
- **Analysis:** For each portfolio, view pie and bar charts of your holdings, and see risk and Sharpe ratio gauges. Click on any asset to view detailed information and price history.

---

## Bugs
There is a bug in the portfolio section, if one of the stocks that the search menu proposes you is not provided
by the financial api, this causes your backend to crash, needing for a restart of the http server.
Unfortunately this could be caused by a change in tickers, which a hardcoded stock ticker list like the one
i use in this project can't keep up.
Using well known stocks is the best choice to avoid this bug.

---

## Disclaimer

This project is for educational purposes only.  
It is not intended for production use or for making real investment decisions.  
Data may be incomplete or delayed, and features may be experimental.

---

## Author

Created by Alessandro / Mister
