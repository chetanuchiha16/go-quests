package main

/// add a concurrency with context for cancellation
import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	var wg sync.WaitGroup

	defer cancel()
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		select {
		case <-time.After(4 * time.Second):
			fmt.Println("done within time")
		case <-ctx.Done():
			fmt.Println("time out")
		}
	}(ctx, &wg)
	wg.Wait()
}
