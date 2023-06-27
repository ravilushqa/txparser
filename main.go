package main

import (
	"context"
	"fmt"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ravilushqa/txparser/ethereum"
	"github.com/ravilushqa/txparser/storage"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Println("error during run: ", err)
	}

}

func run(ctx context.Context) error {
	s := storage.NewMemoryStorage()
	c := ethereum.NewClient("https://cloudflare-eth.com")
	p := ethereum.NewParser(s, c)
	go p.ParseBlocks(ctx)

	subscribers := []string{"0xdac17f958d2ee523a2206206994597c13d831ec7", "0x7830c87C02e56AFf27FA8Ab1241711331FA86F43"}

	for _, subscriber := range subscribers {
		p.Subscribe(strings.ToLower(subscriber))
	}

	// demo
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Second):
			fmt.Println("current block: ", p.GetCurrentBlock())

			for _, subscriber := range subscribers {
				fmt.Println("subscriber: ", subscriber)
				fmt.Println("transactions: ", p.GetTransactions(subscriber))
			}
		}
	}
}
