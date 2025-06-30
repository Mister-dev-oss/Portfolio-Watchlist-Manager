import React from 'react';
import WelcomeBox from '../components/WelcomeBox';
import "./Homepage.css"
import PortfolioBox from '../components/PortfolioBox';
import WatchlistBox from '../components/WatchlistBox';
import AvailableAssetsSearch from '../components/SearchTickerWatchlist';



function HomePage() {
  return (
    <div className='homepage-container'>
      <WelcomeBox />
      <div className='grid2'>
        <PortfolioBox />
        <WatchlistBox />
      </div>
    </div>
  );
}

export default HomePage;

