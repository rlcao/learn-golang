package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup

	wg.Add(2)


	go func() {
		defer wg.Done()
		time.Sleep(1*time.Second)
		fmt.Println("Goroutine1 finished")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2*time.Second)
		fmt.Println("Goroutine 2 finished")
	}()

	wg.Wait()
	fmt.Println("All Jobs Finished")
}
