import React, { useState, useEffect } from 'react';
import Modal from 'react-modal';
import './BtcGbtcModal.css';

const BtcGbtcModal = ({ modalIsOpen, setModalIsOpen, amount, fees, ethAddress, setEthAddress, txHash, setTxHash, signedTxHash, setSignedTxHash, handleBridge }) => {
  const [btcPrice, setBtcPrice] = useState(0);
  const [usdValue, setUsdValue] = useState(0);
  const [loading,setLoading] = useState(false);
  useEffect(() => {
    const fetchBtcPrice = async () => {
      try {
        const response = await fetch('https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd');
        const data = await response.json();
        setBtcPrice(data.bitcoin.usd);
      } catch (error) {
        console.error('Error fetching BTC price:', error);
      }
    };

    fetchBtcPrice();
  }, []);

  useEffect(() => {
    setUsdValue(amount * btcPrice);
  }, [amount, btcPrice]);

  const formatUsdValue = (value) => {
    if (value >= 1000000) {
      return `${(value / 1000000).toFixed(2)}M USD`;
    }
    return `${Math.round(value).toLocaleString()} USD`;
  };

  return (
    <Modal
      isOpen={modalIsOpen}
      onRequestClose={() => setModalIsOpen(false)}
      contentLabel="Bridge Modal"
      style={{
        content: {
          borderRadius: '16px',
          padding: '40px',
          maxWidth: '500px',
          margin: '0 auto',
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
          transition: 'transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out',
        },
        overlay: {
          backgroundColor: 'rgba(0, 0, 0, 0.5)',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }
      }}
    >
      <div className="modal-form-container">
        <h2>Bridge BTC to GBTC</h2>
        <div className="bridge-details">
          <div className="detail-row">
            <span className="detail-title">Amount in BTC:</span>
            <span className="detail-value">{amount}</span>
          </div>
          <div className="detail-row">
            <span className="detail-title">USD Value:</span>
            <span className="detail-value">{formatUsdValue(usdValue)}</span>
          </div>
          <div className="detail-row">
            <span className="detail-title">Fee:</span>
            <span className="detail-value">{fees.toFixed(8)} BTC</span>
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">ETH Destination Address</div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={ethAddress}
                onChange={(e) => setEthAddress(e.target.value)}
                placeholder="Enter your ETH address"
                className="amount-input2"
              />
            </div>
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">Transaction Hash</div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={txHash}
                onChange={(e) => setTxHash(e.target.value)}
                placeholder="Enter the transaction hash"
                className="amount-input2"
              />
            </div>
          </div>
          <div className="bridge-section">
            <div className="bridge-section-header">Signed Transaction Hash</div>
            <div className="bridge-section-content">
              <input
                type="text"
                value={signedTxHash}
                onChange={(e) => setSignedTxHash(e.target.value)}
                placeholder="Enter the signed transaction hash"
                className="amount-input2"
              />
            </div>
          </div>
          <button onClick={()=>setLoading(true)} className="connect-wallet-button">Bridge</button>
          <button onClick={() => setModalIsOpen(false)} className="connect-wallet-button cancel-button">Cancel</button>
        </div>
      </div>
    </Modal>
  );
};

export default BtcGbtcModal;
