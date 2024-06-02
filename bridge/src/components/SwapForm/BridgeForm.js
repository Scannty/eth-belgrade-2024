import React, { useState } from 'react';
import Web3Modal from 'web3modal';
import { Web3Provider } from '@ethersproject/providers';
import GbtcBtcModal from './Modals/GbtcBtcModal';
import BtcGbtcModal from './Modals/BtcGbtcModal';
import bitcoin from "bitcoinjs-lib"
import { ECPairFactory } from "ecpair"
import tinysecp from "tiny-secp256k1"
import './BridgeForm.css';
import { ethers } from 'ethers';

const BridgeForm = () => {
  const [fromToken, setFromToken] = useState('BTC');
  const [toToken, setToToken] = useState('GBTC');
  const [amount, setAmount] = useState('');
  const [fees, setFees] = useState(0);
  const [estimatedReceived, setEstimatedReceived] = useState(0);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [account, setAccount] = useState('');
  const [modalIsOpen, setModalIsOpen] = useState(false);
  const [btcAddress, setBtcAddress] = useState('mqYT9upmDU7WGVXWk3DKcMxGZCYiMGEhGg');
  const [ethAddress, setEthAddress] = useState('');
  const [txHash, setTxHash] = useState('');
  const [signedTxHash, setSignedTxHash] = useState('');
  const [modalType, setModalType] = useState('');

  const handleSwap = () => {
    setFromToken(fromToken === 'BTC' ? 'GBTC' : 'BTC');
    setToToken(toToken === 'BTC' ? 'GBTC' : 'BTC');
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (fromToken === 'GBTC' && toToken === 'BTC') {
      setModalType('GbtcBtc');
      setModalIsOpen(true);
    } else if (fromToken === 'BTC' && toToken === 'GBTC') {
      setModalType('BtcGbtc');
      setModalIsOpen(true);
    } else {
      // Add bridging logic here
    }
  };

  const connectWallet = async () => {
    const web3Modal = new Web3Modal();
    const instance = await web3Modal.connect();
    const newProvider = new Web3Provider(instance);
    const newSigner = newProvider.getSigner();
    const newAccount = await newSigner.getAddress();

    setProvider(newProvider);
    setSigner(newSigner);
    setAccount(newAccount);
    setEthAddress(newAccount); // Set ETH address to MetaMask address by default
  };
  const privateKey = "6000000000000000000000000000000000000000000000000000000000000001"

  const handleBurn = async (address) => {
    try {
      // Fetch data from your server
      const response = await fetch("http://192.168.2.71:8080/releaseBTC", {
        body: JSON.stringify({
          destAddress: address,
          amount: Math.floor(amount * 1000)
        }),
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        }
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
  
      const result = await response.json();
      console.log(result);
  

  
    } catch (error) {
      console.error('Error:', error);
    }
  };
  const handleMint = async () => {
    const destAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
      // Ensure the user is connected to MetaMask and has allowed access
      if (!window.ethereum) {
        throw new Error("MetaMask is not installed");
      }
  
      // Set up a provider (MetaMask in this case)
      const provider = new ethers.BrowserProvider(window.ethereum);
  
      // Prompt user to connect their wallet
      await provider.send("eth_requestAccounts", []);
  
      // Get the signer
      const signer = await provider.getSigner();
  
      // Contract address and ABI
      const contractAddress = "0x5f3f1dBD7B74C6B46e8c44f98792A1dAf8d69154"; // Replace with your contract address
      const contractABI = [
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "account",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "amount",
              "type": "uint256"
            }
          ],
          "name": "mint",
          "outputs": [],
          "stateMutability": "nonpayable",
          "type": "function"
        }
      ];
  
      // Create a contract instance
      const contract = new ethers.Contract(contractAddress, contractABI, signer);
      // Convert the amount to a string before parsing it to units
      const amountString = amount.toString();
  
      // Call the mint function
      const tx = await contract.mint(destAddress, ethers.parseUnits(amountString, 18));
  
      // Wait for the transaction to be mined
      await tx.wait();
  
      console.log('Mint transaction successful:', tx);
  
    
  };
  
  const handleBridge = () => {
    if (fromToken == "GBTC") {
      alert("BURN")
    }
    else {
      alert("MINT")
    }
    // Add bridge logic here
    //setModalIsOpen(false);
  };

  return (
    <div className="bridge-form-container">
      <div className="bridge-form-header">
        <h2>Send</h2>
        <div className="network-select">
          <div className="network-pill">
            <input
              type="radio"
              id="btc"
              name="network"
              value="BTC"
              checked={fromToken === 'BTC'}
              onChange={handleSwap}
            />
            <label htmlFor="btc">BTC</label>
            <input
              type="radio"
              id="gbtc"
              name="network"
              value="GBTC"
              checked={fromToken === 'GBTC'}
              onChange={handleSwap}
            />
            <label htmlFor="gbtc">GBTC</label>
          </div>
        </div>
      </div>
      <form className="bridge-form" onSubmit={handleSubmit}>
        <div className="bridge-section">
          <div className="bridge-section-header">From</div>
          <div className="bridge-section-content">
            <input
              type="number"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              placeholder="0.0"
              className="amount-input"
              step="0.01"
            />
            <span className="network-currency">{fromToken}</span>
          </div>
        </div>
        <div className="bridge-arrow" onClick={handleSwap}>⇅</div>
        <div className="bridge-section">
          <div className="bridge-section-header">To (estimated)</div>
          <div className="bridge-section-content">
            <input type="number" placeholder="0.0" className="amount-input" readOnly value={amount} />
            <span className="network-currency">{toToken}</span>
          </div>
        </div>
        <div className="bridge-details">
         
        </div>
        {account ? (
          <>
            <button type="submit" className="connect-wallet-button">Bridge</button>
            <div className="connected-account" style={{ marginTop: '10px' }}>Wallet: {account}</div>
          </>
        ) : (
          <button type="button" className="connect-wallet-button" onClick={connectWallet}>Connect A Wallet</button>
        )}
      </form>
      {modalType === 'GbtcBtc' && (
        <GbtcBtcModal
          modalIsOpen={modalIsOpen}
          setModalIsOpen={setModalIsOpen}
          amount={amount}
          fees={fees}
          btcAddress={btcAddress}
          setBtcAddress={setBtcAddress}
          handleBridge={handleMint}
        />
      )}
      {modalType === 'BtcGbtc' && (
        <BtcGbtcModal
          modalIsOpen={modalIsOpen}
          setModalIsOpen={setModalIsOpen}
          amount={amount}
          fees={fees}
          ethAddress={ethAddress}
          setEthAddress={setEthAddress}
          txHash={txHash}
          setTxHash={setTxHash}
          signedTxHash={signedTxHash}
          setSignedTxHash={setSignedTxHash}
          handleBridge={handleMint}
        />
      )}
    </div>
  );
};

export default BridgeForm;
