package storage

import "sync"

type MemoryStorage struct {
	mu           sync.Mutex
	subscribers  map[string]bool
	transactions map[string][]Transaction
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		subscribers:  make(map[string]bool),
		transactions: make(map[string][]Transaction),
	}
}

func (s *MemoryStorage) Subscribe(address string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.subscribers[address]
	if exists {
		return false
	}

	s.subscribers[address] = true
	return true
}

func (s *MemoryStorage) GetTransactions(address string) []Transaction {
	s.mu.Lock()
	defer s.mu.Unlock()

	transactions := s.transactions[address]
	s.transactions[address] = []Transaction{} // notify only new transactions

	return transactions
}

func (s *MemoryStorage) SetTransactions(address string, transactions []Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.transactions[address] = transactions
}

func (s *MemoryStorage) GetSubscribers() map[string]bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.subscribers
}
