package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create a buffered channel to hold up to 5 jobs
	jobCh := make(chan int, 5)

	// create a context that cancels after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// spin up 3 worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, jobCh, ctx)
	}

	// use a ticker to send jobs periodically (every 200ms) without a manual loop or sleep
	ticker := time.NewTicker(time.Millisecond * 200)
	defer ticker.Stop()

	jobId := 1

	for {
		select {
		case <-ctx.Done():
			// if the context times out, close the channel and exit
			close(jobCh)
			fmt.Println("Done")
			return
		case <-ticker.C:
			// every tick, send a new job into the job channel
			fmt.Println("Adding job:", jobId)
			jobCh <- jobId
			jobId++
		}
	}
}

// worker simulates a worker picking up jobs from the jobCh
func worker(workerId int, jobCh chan int, ctx context.Context) {
	for {
		select {
		case id, ok := <-jobCh:
			if !ok {
				// if channel is closed, exit worker
				return
			}
			// print which worker is handling which job
			fmt.Printf("Worker %d: Job %d\n", workerId, id)
		case <-ctx.Done():
			// if context ends, stop worker
			return
		}
	}
}
