package main

import (
	"context"
	"fmt"
	"sync/atomic"

	workerpool "github.com/abcxyz/pkg/workerpool"
)

func main() {
	ctx := context.Background()
	pool := workerpool.New[*workerpool.Void](&workerpool.Config{
		Concurrency: int64(10),
		StopOnError: false,
	})
	successCount := atomic.Int64{}

	for i := 0; i < 99999; i++ {
		if err := pool.Do(ctx, func() (*workerpool.Void, error) {
			successCount.Store(successCount.Add(1))
			return nil, nil
		}); err != nil {
			fmt.Printf("pool worker: %s", err)
		}
	}
	_, err := pool.Done(ctx)
	fmt.Println("Success Count: ", successCount.Load())
	if err != nil {
		fmt.Printf("pool error: %s", err)
	}
}
