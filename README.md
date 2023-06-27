# Ethereum Transaction Parser

This project is a simple Ethereum transaction parser written in Go. It subscribes to Ethereum addresses and lists the transactions associated with them.

## How it Works

The project consists of three primary components:

1. `main.go`: The main application file that starts the transaction parser and demonstrates its capabilities.

2. `storage/memory_storage.go`: An in-memory storage that holds the subscriptions and transactions. It's responsible for managing and returning transactions associated with the subscribers.

3. `ethereum/parser.go` and `ethereum/client.go`: The Ethereum client and parser that connect to an Ethereum RPC endpoint, fetch blocks, parse transactions, and update the storage.

## Prerequisites

- Go (version 1.19 or newer)

## Usage

The parser automatically subscribes to the addresses defined in `main.go`. By default, it subscribes to two addresses, and every 10 seconds, it prints the current block number and the transactions of the subscribed addresses.

To add more addresses, add them to the `subscribers` slice in `main.go`.

```go
subscribers := []string{"0xdac17f958d2ee523a2206206994597c13d831ec7", "0x7830c87C02e56AFf27FA8Ab1241711331FA86F43", "<new_address>"}
```

To change the update interval, modify the time in the `time.After` function in `main.go`.

```go
case <-time.After(10 * time.Second): // Change the interval here
```

The parser uses `https://cloudflare-eth.com` as the Ethereum RPC endpoint. You can change this by modifying the URL in `ethereum.NewClient`.

```go
c := ethereum.NewClient("https://cloudflare-eth.com") // Change the URL here
```

## Important Note

This parser is not production-ready. It's a simple demonstration of how to parse Ethereum transactions using Go. 