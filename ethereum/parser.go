package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/ravilushqa/txparser/storage"
)

var _ Parser = (*EthParser)(nil)

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []storage.Transaction
}

type EthParser struct {
	storage      storage.Storage
	currentBlock atomic.Int64
	client       *Client
}

type Block struct {
	Transactions []storage.Transaction
}

func NewParser(s storage.Storage, c *Client) *EthParser {
	return &EthParser{storage: s, client: c}
}

func (p *EthParser) GetCurrentBlock() int {
	return int(p.currentBlock.Load())
}

func (p *EthParser) Subscribe(address string) bool {
	return p.storage.Subscribe(address)
}

func (p *EthParser) GetTransactions(address string) []storage.Transaction {
	return p.storage.GetTransactions(address)
}

func (p *EthParser) ParseBlocks(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
		}
		blockNumber, err := p.getLatestBlockNumber()
		if err != nil {
			fmt.Println("Error getting latest block number", err)
			continue
		}

		if blockNumber == p.GetCurrentBlock() {
			fmt.Println("No new blocks")
			continue
		}

		fmt.Println("New block", blockNumber)
		block, err := p.getBlockByNumber(blockNumber)
		if err != nil {
			fmt.Println("Error getting block", blockNumber, err)
			continue
		}

		for _, tx := range block.Transactions {
			subs := p.storage.GetSubscribers()
			if _, ok := subs[tx.From]; ok {
				p.storage.SetTransactions(tx.From, append(p.storage.GetTransactions(tx.From), tx))
			}

			if _, ok := subs[tx.To]; ok {
				p.storage.SetTransactions(tx.To, append(p.storage.GetTransactions(tx.To), tx))
			}
		}
		p.currentBlock.Store(int64(blockNumber))
	}
}

func (p *EthParser) getLatestBlockNumber() (int, error) {
	requestBody := RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []string{},
		ID:      1,
	}

	rpcResponse, err := p.client.SendRequest(requestBody)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %w", err)
	}

	var blockHash string
	err = json.Unmarshal(rpcResponse.Result, &blockHash)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// Convert the block hash to an int
	var blockNumber int
	_, err = fmt.Sscanf(blockHash, "0x%x", &blockNumber)
	if err != nil {
		return 0, fmt.Errorf("error converting block hash to int: %w", err)
	}

	return blockNumber, nil
}

func (p *EthParser) getBlockByNumber(blockNumber int) (Block, error) {
	requestBody := RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []string{fmt.Sprintf("0x%x", blockNumber), "true"},
		ID:      1,
	}

	rpcResponse, err := p.client.SendRequest(requestBody)
	if err != nil {
		return Block{}, fmt.Errorf("error sending request: %w", err)
	}

	var block Block
	err = json.Unmarshal(rpcResponse.Result, &block)
	if err != nil {
		return Block{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return block, nil
}
