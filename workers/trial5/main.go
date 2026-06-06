package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	defer close(ch)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go func(ctx context.Context, ch chan bool) {
		select {
		case <-time.After(2 * time.Second):
			ch <- false
		case <-ctx.Done():
			ch <- true
		}
	}(ctx, ch)
	if isFast := <-ch; isFast {
		fmt.Printf("were fast")
	} else {
		fmt.Println("slow")
	}
	fmt.Println()

}
