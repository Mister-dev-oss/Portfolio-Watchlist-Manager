import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import './Sidebar.css';
import { FaHome, FaFolderOpen, FaInfoCircle } from 'react-icons/fa';

const links = [
  { to: '/', icon: <FaHome />, label: 'Home' },
  { to: '/portfolio', icon: <FaFolderOpen />, label: 'Portfolios' },
  { to: '/info', icon: <FaInfoCircle />, label: 'Info' },
];

export default function Sidebar() {
  const location = useLocation();

  return (
    <div className="sidebar">
      <div className="logo-placeholder" />

      <div className="nav-links">
        {links.map((link) => (
          <Link
            key={link.to}
            to={link.to}
            className={`nav-circle ${location.pathname === link.to ? 'active' : ''}`}
            title={link.label}
          >
            {link.icon}
          </Link>
        ))}
      </div>
    </div>
  );
}

