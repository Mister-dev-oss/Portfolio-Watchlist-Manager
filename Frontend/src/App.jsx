import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/Homepage';
import PortfolioPage from './pages/PortfolioPage';
import TickerPage from './pages/TickerPage';
import InfoPage from './pages/InfoPage';
import Sidebar from './components/SideBar';
import './App.css';

function App() {
  return (
    <Router>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <Sidebar />
        <div style={{ marginLeft: '5rem', padding: '0rem', marginTop: '2rem'}}>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/portfolio" element={<PortfolioPage />} />
            <Route path="/ticker" element={<TickerPage />} />
            <Route path="/info" element={<InfoPage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App