const bitcoin = require('bitcoinjs-lib');
const { payments } = bitcoin;
const bs58check = require('bs58check');
const { ECPairFactory } = require("ecpair");
const tinysecp = require('tiny-secp256k1');
const bitcoinMessage = require('bitcoinjs-message');

const ECPair = ECPairFactory(tinysecp);

// Your private key in hex
const privateKeyHex = '1010101010101010101010101010101010101010101010101010101000000001';

// Create a key pair from the private key
const keyPair = ECPair.fromPrivateKey(Buffer.from(privateKeyHex, 'hex'), { network: bitcoin.networks.regtest });

// The message to sign
const messageToSign = 'kurac';

// Generate the address from the key pair
const { address } = payments.p2pkh({ pubkey: keyPair.publicKey, network: bitcoin.networks.regtest });

const privateKeyBuffer = keyPair.privateKey;
// Sign the message
const signature = bitcoinMessage.sign(messageToSign, privateKeyBuffer, keyPair.compressed);

// Convert signature to base64
const signatureBase64 = signature.toString('base64');

console.log('Signed Message:', signatureBase64);


const isValid = bitcoinMessage.verify(messageToSign, "mt3W59U5sdSCn6e8QPfMGjZezE6sAKkZoX", signatureBase64);
console.log('Message:', messageToSign);
console.log('Address:', address);
console.log('Signature is valid:', isValid);


