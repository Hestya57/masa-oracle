# Masa Oracle: Decentralized Data Protocol 🌐

The Masa Oracle governs the access, sharing, and rewarding of private behavioral and identity data in a decentralized and private manner. The Masa Oracle Network ensures transparency and security of data sharing, while  enabling equitable compensation for nodes that participate in the Masa zk-Data Network and Marketplace.

## Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Docker Setup](#docker-setup)
- [Staking Tokens](#staking-tokens)
- [Running the Node](#running-the-node)
- [Updates & Additional Information](#updates--additional-information)

## Getting Started

### Prerequisites

Ensure these prerequisites are installed for a local setup:
- **Go**: Grab it from [Go's official site](https://golang.org/dl/).
- **Yarn**: Install it via [Yarn's official site](https://classic.yarnpkg.com/en/docs/install/).
- **Git**: Required for cloning the repository.
- For complete instructions on building, staking, and running a node with Docker, please see [here](./DOCKER.md) 

### Installation


#### Docker Setup

For complete instructions on building, staking, and running a node with Docker, please see [here](./DOCKER.md) 

#### Local Setup

1. Clone the repository
```
git clone https://github.com/masa-finance/masa-oracle.git
```
2. Build the go code into the masa-node binary:
```
go build -v -o masa-node ./cmd/masa-node
```
3. Go into the contracts directory and build the contract npm modules that the go binary uses:
```
cd contracts/ 
npm install
cd ../
```
4. Set env vars
Ensure your environment has the required env var, ENV, exported. It should be set to test to join the testnet.
Optionally, you can set RPC_URL to change it from the default (https://ethereum-sepolia.publicnode.com)
```
export ENV=test
export RPC_URL=https://1rpc.io/sepolia	# This is optional and will be set to the default https://ethereum-sepolia.publicnode.com without it.
```

5. Start up masa-node. Be sure to include your bootnodes list with the --bootnodes flag.
```
/ip4/34.116.150.247/udp/4001/quic-v1/p2p/16Uiu2HAmHpx13GPKZAP3WpgpYkZ39M5cwuvmXS5gGvrsa5ofLNoq
/ip4/34.118.84.78/udp/4001/quic-v1/p2p/16Uiu2HAm7HAXW7HS1YA4mua8oyo8Se9cTk8MSXXRUmsJqN7NxhB1
/ip4/34.118.116.5/udp/4001/quic-v1/p2p/16Uiu2HAm5zfUVMQrBMDfLaA5xwPXhpwAYzTyrkTsHoqPzmj8PVLz
/ip4/34.116.165.47/udp/4001/quic-v1/p2p/16Uiu2HAmBazcN1AaF2KcvfurwAJdvCJiD8BW6itRRqBsdqvXByb9
/ip4/34.72.224.59/udp/4001/quic-v1/p2p/16Uiu2HAm17obtAHet7YkoPH1vcsteBYFVmNJq62gGEJ5xxSu5BAk
```

```
masa-node masa-node --start --bootnodes=/ip4/35.223.224.220/udp/4001/quic-v1/p2p/16Uiu2HAmPxXXjR1XJEwckh6q1UStheMmGaGe8fyXdeRs3SejadSa
```

## Funding the Node (in order to Stake)


Find the public key of your node in the logs. 

Send 1000 MASA and .01 sepoliaETH to the node's public key / wallet address.

When the transactions have settled, you can stake

### Staking Tokens

- For local setup, stake tokens with:
  ```bash
  ./masa-node --stake 1000
  ```
- For Docker setup, stake tokens with:
  ```bash
  docker-compose run --rm masa-node /usr/bin/masa-node --stake 1000
  ```

### Running the Node

- **Local Setup**: Connect your node to the Masa network:
  ```bash
  ./masa-node --bootnodes=/ip4/35.223.224.220/udp/4001/quic-v1/p2p/16Uiu2HAmPxXXjR1XJEwckh6q1UStheMmGaGe8fyXdeRs3SejadSa --port=4001 --udp=true --tcp=false --start=true --env=test
  ```
- **Docker Setup**: Your node will start automatically with `docker-compose up -d`. Verify it's running correctly:
  ```bash
  docker-compose logs -f masa-node
  ```

After setting up your node, its address will be displayed, indicating it's ready to connect with other Masa nodes. Follow any additional configuration steps and best practices as per your use case or network requirements.

## Updates & Additional Information

Stay tuned to the Masa Oracle repository for updates and additional details on effectively using the protocol. For Docker users, update your node by pulling the latest changes from the Git repository, then rebuild and restart your Docker containers.

