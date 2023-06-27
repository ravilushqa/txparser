package ethereum

import (
	"testing"

	"github.com/ravilushqa/txparser/storage"
)

func TestGetCurrentBlock(t *testing.T) {
	client := NewClient("https://cloudflare-eth.com")
	memoryStorage := storage.NewMemoryStorage()
	parser := NewParser(memoryStorage, client)

	if parser.GetCurrentBlock() != 0 {
		t.Errorf("Initial block is not zero.")
	}

	parser.currentBlock.Store(10)

	if parser.GetCurrentBlock() != 10 {
		t.Errorf("Block value changed is not reflected.")
	}
}

func TestSubscribe(t *testing.T) {
	client := NewClient("https://cloudflare-eth.com")
	memoryStorage := storage.NewMemoryStorage()
	parser := NewParser(memoryStorage, client)

	subscribed := parser.Subscribe("0x0")

	if !subscribed {
		t.Errorf("Subscribe returned false for a new address.")
	}
}

func TestGetTransactions(t *testing.T) {
	client := NewClient("https://cloudflare-eth.com")
	memoryStorage := storage.NewMemoryStorage()
	parser := NewParser(memoryStorage, client)

	transactions := parser.GetTransactions("0x0")

	if len(transactions) != 0 {
		t.Errorf("New address has transactions.")
	}
}
