package main

import "fmt"

func fanout(workers int, tasks <-chan string) []chan string {
	outs := make([]chan string, workers)
	for worker := range workers {
		out := make(chan string)
		go func(worker int) {
			defer close(out)
			for task := range tasks {
				result := fmt.Sprintf("%v by worker %v", task, worker)
				out <- result
			}
		}(worker + 1)
		outs[worker] = out
	}
	return outs
}

func main() {
	workers := 5
	tasks := make(chan string)
	outs := fanout(workers, tasks)

	go func() {
		defer close(tasks)
		for i := range 10 {
			tasks <- fmt.Sprintf("task %v", i+1)
		}
	}()

	for _, out := range outs {
		for o := range out {

			fmt.Println(o)
		}
	}
}
