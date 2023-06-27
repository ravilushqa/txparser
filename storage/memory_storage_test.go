package storage

import (
	"reflect"
	"testing"
)

func TestMemoryStorage_Subscribe(t *testing.T) {
	store := NewMemoryStorage()

	if store.Subscribe("123") != true {
		t.Errorf("Expected true when subscribing an unused address, but got false.")
	}

	if store.Subscribe("123") != false {
		t.Errorf("Expected false when subscribing an already used address, but got true.")
	}
}

func TestMemoryStorage_GetTransactions(t *testing.T) {
	store := NewMemoryStorage()
	store.SetTransactions("123", []Transaction{
		{
			From:  "123",
			To:    "456",
			Value: "10",
			Hash:  "hash1",
		},
		{
			From:  "123",
			To:    "789",
			Value: "20",
			Hash:  "hash2",
		},
	})

	transactions := store.GetTransactions("123")

	if len(transactions) != 2 {
		t.Errorf("Expected two transactions, but got %d", len(transactions))
	}
}

func TestMemoryStorage_SetTransactions(t *testing.T) {
	store := NewMemoryStorage()
	store.SetTransactions("123", []Transaction{
		{
			From:  "123",
			To:    "456",
			Value: "10",
			Hash:  "hash1",
		},
		{
			From:  "123",
			To:    "789",
			Value: "20",
			Hash:  "hash2",
		},
	})

	transactions := store.GetTransactions("123")

	if len(transactions) != 2 {
		t.Errorf("Expected two transactions, but got %d", len(transactions))
	}
}

func TestMemoryStorage_GetSubscribers(t *testing.T) {
	store := NewMemoryStorage()
	store.Subscribe("123")
	store.Subscribe("456")

	subscribers := store.GetSubscribers()

	expectedSubscribers := map[string]bool{
		"123": true,
		"456": true,
	}

	if !reflect.DeepEqual(subscribers, expectedSubscribers) {
		t.Errorf("Expected subscribers to be %v, but got %v", expectedSubscribers, subscribers)
	}
}
