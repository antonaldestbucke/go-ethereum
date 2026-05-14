# go-ethereum

Official Go implementation of the Ethereum protocol — this repository is a fork of [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum).

[![API Reference](https://pkg.go.dev/badge/github.com/ethereum/go-ethereum)](https://pkg.go.dev/github.com/ethereum/go-ethereum?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethereum/go-ethereum)](https://goreportcard.com/report/github.com/ethereum/go-ethereum)
[![Travis](https://travis-ci.com/ethereum/go-ethereum.svg?branch=master)](https://travis-ci.com/ethereum/go-ethereum)
[![Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.gg/nthXNEv)

## About

This is a fork of the official Ethereum Go client (`geth`). It includes all upstream changes along with additional patches, performance improvements, and custom configurations suited for our deployment needs.

## Building the source

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/getting-started/installing-geth).

Building `geth` requires both a Go (version 1.21 or later) and a C compiler. You can install them using your favourite package manager. Once the dependencies are installed, run:

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-ethereum project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
| :--------- | :---------- |
| **`geth`** | Main Ethereum CLI client. It is the entry point into the Ethereum network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ethereum network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. |
| `clef`     | Stand-alone signing tool, which can be used as a backend signer for `geth`. |
| `devp2p`   | Utilities to interact with nodes on the networking layer, without running a full blockchain. |
| `abigen`   | Source code generator to convert Ethereum contract definitions into easy-to-use, compile-time type-safe Go packages. |
| `bootnode` | Stripped down version of our Ethereum client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. |
| `evm`      | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. |
| `rlpdump`  | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation. |

## Running `geth`

### Hardware Requirements

Minimum:

* CPU with 2+ cores
* 4GB RAM
* 1TB free storage space to sync the Mainnet
* 8 MBit/sec download Internet service

Recommended:

* Fast CPU with 4+ cores
* 16GB+ RAM
* High-performance SSD with at least 1TB of free space
* 25+ MBit/sec download Internet service

### Full node on the main Ethereum network

By far the most common scenario is people wanting to simply interact with the Ethereum network: create accounts; transfer funds; deploy and interact with contracts. For this particular use case, the user doesn't care about years-old historical data, so we can sync quickly to the current state of the network. To do so:

```shell
geth console
```

## Contributing

Thank you for considering contributing to the go-ethereum source code! Please see the [CONTRIBUTING.md](.github/CONTRIBUTING.md) file for details on our code of conduct and the process for submitting pull requests.

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) are licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `COPYING` file.
