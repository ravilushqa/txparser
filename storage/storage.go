package storage

type Storage interface {
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
	SetTransactions(address string, transactions []Transaction)
	GetSubscribers() map[string]bool
}

type Transaction struct {
	From  string
	To    string
	Value string
	Hash  string
}
