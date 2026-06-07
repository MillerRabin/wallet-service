# Wallet Service

Wallet Service is a Go application responsible for HD wallet operations.

The service derives addresses and private keys from a master mnemonic using BIP39/BIP44 standards and signs blockchain transactions.

Private keys are never exposed through the API.

The service is designed to work together with Gates:

https://github.com/MillerRabin/gates

Gates is responsible for blockchain indexing, deposit detection, withdrawal orchestration and transaction broadcasting, while Wallet Service is responsible exclusively for key derivation and transaction signing.

---

# Architecture

```text
+----------------+
|     Gates      |
|   Laravel API  |
+--------+-------+
         |
         | HTTP
         v
+----------------+
| Wallet Service |
|      Go        |
+--------+-------+
         |
         | Signed TX
         v
+----------------+
| Ethereum RPC   |
+----------------+
```

Components:

* HTTP API
* HD Wallet Engine
* Ethereum Signer

The service uses:

* BIP39
* BIP32
* BIP44
* secp256k1
* EIP-1559 transactions

No database is required.

All addresses and private keys are deterministically derived from a single mnemonic.

---

# Security Model

The system is split into two independent services.

Wallet Service owns:

* mnemonic phrase
* private keys
* HD wallet derivation
* transaction signing

Gates owns:

* deposits
* withdrawals
* blockchain indexing
* transaction broadcasting

Private keys never leave Wallet Service.

A compromise of Gates does not expose private keys.

---

# Trust Wallet Core

Wallet Service is built on top of Trust Wallet Core:

https://github.com/trustwallet/wallet-core

Trust Wallet Core provides:

* BIP39 mnemonic support
* BIP32/BIP44 HD wallet derivation
* Ethereum transaction signing
* Multi-chain cryptographic primitives

Trust Wallet Core is built using the official Docker build process.

All required libraries, generated protobuf definitions and dependencies are prepared automatically through the Makefile.

No manual installation of Trust Wallet Core is required.

---

# Supported Networks

Currently supported:

* Ethereum
* Ethereum Sepolia

Additional EVM-compatible networks can be added by providing the appropriate chain ID.

---

# HD Wallet Path

Ethereum addresses are derived using:

```text
m/44'/60'/account'/change/address_index
```

Example:

```text
m/44'/60'/0'/0/0
```

Parameters:

* account
* change
* address_index

Example request:

```json
{
  "account": 0,
  "change": 0,
  "address_index": 15
}
```

---

# Configuration

Before starting the application, create a working configuration file:

```bash
cp config.example.json config.json
```

Edit `config.json` and replace the example mnemonic with your own mnemonic phrase.

Example:

```json
{
  "config": {
    "host": "0.0.0.0",
    "port": 8000
  },
  "gates": [
    {
      "name": "ethereum",
      "mnemonic": "your mnemonic phrase here"
    }
  ]
}
```

The mnemonic is the master seed used for:

* HD wallet derivation
* address generation
* transaction signing

The `config.json` file contains sensitive information and should never be committed to version control.

Recommended:

```gitignore
config.json
```

The repository includes `config.example.json` as a template for local development.

---

# Quick Start

Clone repository:

```bash
git clone https://github.com/MillerRabin/wallet-service.git
```

Install dependencies:

```bash
make install
```

Build application:

```bash
make build
```

Run application:

```bash
make run
```

---

# Development

Format code:

```bash
make fmt
```

Run static analysis:

```bash
make vet
```

Update Go modules:

```bash
make tidy
```

---

# Testing

Run tests:

```bash
make test
```

Run race detector:

```bash
make test-race
```

---

# Docker

Build Docker image:

```bash
make docker-build
```

Run container:

```bash
make docker-run
```

Stop container:

```bash
make docker-stop
```

---

# API

Base URL:

```text
http://localhost:8000/api/v1
```

---

## Create Address

POST /api/v1/createaddress

Request:

```json
{
  "gate": "ethereum",
  "account": 0,
  "change": 0,
  "address_index": 0
}
```

Response:

```json
{
  "address": "0x9858EfFD232B4033E47d90003D41EC34EcaEda94"
}
```

---

## Validate Address

POST /api/v1/validateaddress

Request:

```json
{
  "gate": "ethereum",
  "address": "0x9858EfFD232B4033E47d90003D41EC34EcaEda94"
}
```

Response:

```json
{
  "valid": true
}
```

---

## Create Transaction

POST /api/v1/tx

Request:

```json
{
  "gate": "ethereum",
  "account": 0,
  "change": 0,
  "address_index": 0,
  "tx_params": {
    "to": "0x2222222222222222222222222222222222222222",
    "value_wei": "400000000000000",
    "data": "0x",
    "nonce": 1,
    "chain_id": 11155111,
    "gas_limit": 21000,
    "max_fee_per_gas_wei": "30000000000",
    "max_priority_fee_per_gas_wei": "1500000000"
  }
}
```

Response:

```json
{
  "tx_hash": "0x...",
  "signed_tx": "0x..."
}
```

---

# ERC20 Transactions

For ERC20 transfers:

* to = token contract address
* value_wei = 0
* data = encoded transfer(address,uint256)

Example:

```text
transfer(
    recipient,
    amount
)
```

The service signs transactions but does not broadcast them.

Transaction broadcasting is handled by Gates.

---

# Related Projects

Gates

https://github.com/MillerRabin/gates

Trust Wallet Core

https://github.com/trustwallet/wallet-core
