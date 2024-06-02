import React from 'react';
import './Dashboard.css';
import BridgeForm from '../SwapForm/BridgeForm';
import logo from '../../assets/logo.png';

const Dashboard = () => {
  return (
    <div className="dashboard-container">
      <div className="dashboard-text">
        <img src={logo} alt="Logo" className="logo" />
        <h1>Gazelle, Trustless Bitcoin-Ethereum Bridge</h1>
        <p>
          This project aims to provide a trustless, decentralized bridge between Bitcoin and Ethereum networks. By utilizing Eigen Layer and wrapped tokens, it allows for seamless, secure swaps between BTC and GBTC, ensuring that users can leverage the advantages of both blockchains without relying on a central authority.
        </p>
        <p>
          <a href="https://github.com/your-repo-link" target="_blank" rel="noopener noreferrer">
            View the Code Repository
          </a>
        </p>
        <p>
          <a href="https://eigenlayer.xyz" target="_blank" rel="noopener noreferrer" className="eigen-layer-button">
            Find out more about Eigen Layer
          </a>
        </p>
      </div>
      <div className="dashboard-swap">
        <BridgeForm />
      </div>
    </div>
  );
};

export default Dashboard;

