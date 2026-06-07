package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func withBuffered() {
	fmt.Printf("Goroutines at start with no buffer: %d\n", runtime.NumGoroutine())
	ch := make(chan bool, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func(ctx context.Context, ch chan bool) {
		select {
		case <-time.After(1 * time.Second):
			ch <- false
		case <-ctx.Done():
			ch <- true
		}
		fmt.Println("exiting go routine")
	}(ctx, ch)

	if patient := false; patient {
		if isFast := <-ch; isFast {
			fmt.Printf("were fast")
		} else {
			fmt.Println("slow")
		}
	} else {
		fmt.Println("im impatient")
	}

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Goroutines at the end: %d\n", runtime.NumGoroutine())
}
func withoutBuffered() {
	fmt.Printf("Goroutines at start with no buffer: %d\n", runtime.NumGoroutine())
	
	ch := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func(ctx context.Context, ch chan bool) {
		select {
		case <-time.After(1 * time.Second):
			ch <- false
		case <-ctx.Done():
			ch <- true
		}
		fmt.Println("exiting go routine")
	}(ctx, ch)

	if patient := false; patient {
		if isFast := <-ch; isFast {
			fmt.Printf("were fast")
		} else {
			fmt.Println("slow")
		}
	} else {
		fmt.Println("im impatient")
	}

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Goroutines at the end: %d\n", runtime.NumGoroutine())
}
func main() {
	fmt.Println("running without buffered channel")
	withoutBuffered()
	
	fmt.Println("\n\nrunning with buffered channel")
	withBuffered()
}
