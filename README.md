
Motif Node is the core infrastructure component of the Motif Protocol — a decentralized platform enabling Bitcoin liquidity providers to create Distributed Traded Products (DTPs). It facilitates the seamless locking of Bitcoin through Bitcoin Pods, which is secured by EigenLayer shared services.

The node is responsible for:
	•	Managing Bitcoin deposits and their representation on Ethereum.
	•	Interacting with EigenLayer for shared security.
	•	Supporting vault strategies that generate yield through integrated DeFi protocols.

This repository contains the core logic for running a Motif Node, including smart contract interactions, cross-chain operations, and node management tools.


# Motif Node Setup Guide 🧪 ⚙️

## 1. Overview
The architecture includes the following components:
- **Event Manager**: A thread which registers to the events of `Motif` AVS contract.
- **Deposit Checker**: A thread which confirms the BTC deposit form a BTC node . 
- **Api Server**: An API server to create new multisig addresses and PSBTs. 
- **bitcoind Offline Wallet**: An offline server that hosts a single wallet containing a signer BTC key. This server is used to sign PSBT transactions forwarded by the `API server`.
- **Signet Node**: A `BTC` full `Signet` node with txindexing enabled, used to watch addresses, create transactions and broadcast them. 


**For a production environment, it is highly recommended to deploy the bitcoind Offline Wallet and the nyks Full Node/btc-oracle on separate hosts**.

The btc-oracle offline signer design is based on remote signing available in `bitcoind` and `lnd`. Signer mode, however, does not require a BTC full node connection as the signer is not responsible for creating or broadcasting transactions.

References:
- [Bitcoind: Managing the Wallet](github.com/bitcoin/bitcoin/blob/master/doc/managing-wallets.md)
- [Bitcoind: Offline Signing Tutorial](github.com/bitcoin/bitcoin/blob/master/doc/offline-signing-tutorial.md)
- [Lightning: Remote Signing](github.com/lightningnetwork/lnd/blob/master/docs/remote-signing.md)

## 2. BTC offline wallet setup

We recommend using an instance of [bitcoin-core] (https://bitcoin.org/en/releases/27.0/) configured in the Offline signing wallet mode.The Bitcoin Core wallet is the preferred choice because it enables clients to utilize external signers and boasts a long-standing, rigorously tested codebase.

For ensuring wallet security, it is strongly recommended to use a separate host system with atleast 4 GB RAM and 2 GB available storage space. The system should be completely disconnected from all public networks (internet, tor, wifi etc.). The `offline` wallet host is not required to download or synchronize blockchain data.

The offline wallet should be setup as a server as part of a secured private network where, the wallet is only accessible through a designated rpc connection with `btc-oracle`. To ensure the integrity and security of the data, only `TLS` based rpc connection should be allowed between `btc-oracle` and the offline wallet node. 

### 2.1 Installation

Download and install the bitcoin binaries according to your operating systemfrom the official [Bitcoind Core registry](https://bitcoincore.org/bin/bitcoin-core-27.0/). All programs in this guide are compatible with version `27.0`.

### 2.2 Configuration
`bitcoind` instance can be configured by using a `bitcoin.conf` file. In `Linux` based systems the file is found in `/home/<username>/.bitcoin`.

A sample configuration file with recommended settings is as follows
```shell
# Accept command line and JSON-RPC commands
server=1

# RPC server settings
rpcuser=<rpc-username>
rpcpassword=<rpc-password>
# field <userpw> comes in the format: <USERNAME>:<SALT>$<HASH>.
# rpcauth = <userpw>

# Port your bitcoin node will listen for incoming requests;
# listening for bitcoin mainnet
rpcport=8332 
# Address your bitcoin node will listen for incoming requests
# should be the address of your offline host
rpcbind=0.0.0.0
# Needed for remote node connectivity
# btc-oracle IP should only be allowed 
rpcallowip=0.0.0.0/0
# Offline Wallet server shouldn't connect to any external p2p or chain node
connect=0
```

JSON-RPC connection authentication can be configured to use `rpc-username`:`rpc-password` pair or a `username and HMAC-SHA-256 hashed password` through rpcauth option. It is not recommended to hardcode `rpc-password` in the config file. The salted hash can be created from canonical python script included in the share/rpcauth in bitcoin-core installed directory. 

### 2.3. Run the RPC Server 

By default, the `bitcoind` server can be run using the following command.
```shell
bitcoind
```
In case, a non-default home directory was used during installation:

```shell
bitcoind -datadir=/path/to/bitcoin/home
```
### 2.3. Manage Wallet 

The following commands shall be used to create and manage BTC wallet on the offline host. The wallet will contain a single address controlled by a private key that will be used to sign transactions recieved from `btc-oracle`

1. Create the wallet
    ```shell
    bitcoin-cli -named createwallet \
        wallet_name=<wallet_name> \
        passphrase=<passphrase> \
        load_on_startup=true \
    ```
    Flags explanation:
    - `wallet_name`: The name of the wallet
    - `passphrase`: The passphrase that will be used to encrypt the `wallet.dat` file. 
    - `load_on_startup=true`: Ensures that the wallet is automatically loaded in
      case of `bitcoind` server restart

2. Create a new address
    ```shell
    bitcoin-cli getnewaddress
    ```
    Save the address for future use. This address shall be used to recieve/send funds on the BTC chain. 

3. Obtain 33-byte BTC public key derived from the above address
    ```shell
    bitcoin-cli getaddressinfo <btc_address> | jq -r .pubkey
    ```
    Maintain a record of the public key in its hexadecimal string fromat. The btc public key will be used during the signer registration process when setting up the `btc-oracle` to act as Fragment `Signer`.  

4. The wallet can be unlocked on the offline host using the following command
    ```shell
      bitcoin-cli walletpassphrase <passphrase> <unlock_time>
    ```
    where:
    - `passphrase` is the same as when used for creating the wallet and,
    - `unlock_time` is the amount of time the wallet is unlocked for in seconds

5. The wallet can be backedup and restored 
    ```shell
      # Backup the wallet
      bitcoin-cli -rpcwallet=<wallet-name> backupwallet /path/to/backup/wallet.dat
      # Restore the wallet
      bitcoin-cli restorewallet <wallet-name> /path/to/backup/wallet.dat
    ```
    It is recommended to take periodic backups of the wallet to keep it secure.     


## 3. Configuration file
The config/config.json file contains various configuration settings required for connecting to Bitcoin and Ethereum nodes, as well as other related services. Below is a detailed explanation of each configuration parameter:

### 3.1 Bitcoin Configuration
- `btc_node_pass`: The password for accessing the Bitcoin node.
- `btc_node_protocol`: The protocol used to connect to the Bitcoin node (e.g., http://).
- `fee_rate_adjustment`: The adjustment rate for transaction fees. Increses the btc tx fee per byte (in sats)
- `wallet_name`: The name of the Bitcoin wallet used for transactions.
- `btc_node_host`: The host URL for connecting to the Bitcoin node.
- `btc_node_user`: The username for accessing the Bitcoin node.
- `btc_xpublic_key`: The extended public key for the Bitcoin wallet, used for generating addresses.
- `multisig_signing_wallet_name`: Offline wallet name,
- `multisig_btc_node`: Offline btc node,
- `multisig_btc_user`: "Offline Btc user",
- `multisig_btc_pass`: "Offline Btc Pass",
- `multisig_btc_protocol`: "Offline Btc protocol,

### 3.2 Ethereum Configuration
- `eth_rpc_host`: The RPC host URL for connecting to the Ethereum node.
- `eth_ws_host`: The WebSocket host URL for connecting to the Ethereum node.
- `eth_keystore_dir`: The directory where the Ethereum keystore files are stored.
- `eth_keystore_passphrase`: The password for accessing the Ethereum keystore.
### 3.3 AVS Configuration
- `opr_metadata_uri`: The URI for the OPR metadata (currently empty).
- `eigen_delegation_manager_address`: The Ethereum address of the Eigen Delegation Manager contract.
- `motif_registry_address`: The Ethereum address of the Motif Registry contract.
- `service_manager_address`: The Ethereum address of the Service Manager contract.
- `eigen_avs_directory_address`: The Ethereum address of the Eigen AVS Directory contract.
- `pod_manager_address`: The Ethereum address of the Pod Manager contract
### 3.4 Database Configuration
- `DB_port`: The port number for the database connection.
- `DB_host`: The host URL for the database connection.
- `DB_user`: The username for the database connection.
- `DB_password`: The password for the database connection.
- `DB_name`: The Name of the database.

This configuration file should be updated with the appropriate values for your specific setup. Ensure that sensitive information such as passwords and private keys are securely managed.

### 3.5. Contract addresses
The contract addresses for the deployed contracts are available on the [Motif repo.](https://github.com/MotifFinance/motif-node) 


## 4. Eth Wallet Setup
The System Uses the go-ethereum library's keystore save keys. if there is no keystore file in the directory (mentioned in config file) the system will generate a new key and save it in a newly created keystore located at the mentioned directory encrypted with the password mentioned in config file. 

If you have a Private key you want to use specifically, you can use the code snippet below to create a keystore before you run Motif Operator.

```go
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <private_key_hex> <password>")
		return
	}

	privateKeyHex := os.Args[1]
	password := os.Args[2]

	// Decode the private key from hex
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	// Convert the private key bytes to an ECDSA private key
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	// Create a keystore
	ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// Create a new account with the private key
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		log.Fatalf("Failed to import private key to keystore: %v", err)
	}

	fmt.Printf("Keystore created: %s\n", account.Address.Hex())
}
```

## 5. Run the Operator

### 5.1. Install Golang
To run the operator, you will first have to install [Golang](https://go.dev/doc/install)

### 5.2. Build and Runthe Operator
To run the operator run the following commands from root

```shell
go mod tidy 
go build .
go run main.go
```

This will start all the components of the operator. It is ideal to run it inside a process manager like `systemd` or `tmux` to ensure that the operator is running continuously. 


## 6. API Server
The API server is used to create new multisig addresses and PSBTs. It is a simple HTTP server that listens on port 8080. 

### 6.1. Create a new multisig address
- `Endpoint`: /eigen/get_address
- `Method`: POST
- `Description`: It is a POST operation that returns a new multisig generated using the provided pubkey and the pubkey of the operator.
- `Body` :
```json
{
    "pubKey":"",
}
```
### 6.2 Api endpoints to HexToBech32 conversion and vice versa
- `Endpoint`: /eigen/bech32_to_hex ans /eigen/hex_to_bech32
- `Method`: POST
- `Description`: It is an endpoint used to convert bech32 to a hex string and vice versa
- `Body` :
```json
{
    "address":"",
}
```

### 6.3. Get Node Information
- `Endpoint`: /eigen/node
- `Method`: GET
- `Description`: Retrieves information about the node.

### 6.4. Health Check
- `Endpoint`: /eigen/node/health
- `Method`: GET
- `Description`: Checks the health status of the node.

### 6.5. Get Services
- `Endpoint`: /eigen/node/services
- `Method`: GET
- `Description`: Retrieves a list of services running on the node.

### 6.6. Get Service Health
- `Endpoint`: /eigen/node/services/{service_ID}/health
- `Method`: GET
- `Description`: Checks the health status of a specific service identified by service_ID.


## 7. DataBase Setup
We use Postgres DB. Once the DB is installed you can run the .db/schema.sql file to create the tables.
