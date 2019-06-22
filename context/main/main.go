package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	ctx,cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Job is over!")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1* time.Second)
			}
		}
	}(ctx)

	time.Sleep(3*time.Second)
	fmt.Println("Time is up, stopping")
	cancel()
	time.Sleep(1*time.Second)
	fmt.Println("Everything is over!")
}

