package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	fmt.Println("Start Goroutines ...")

	//waitGroup
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for c:= 'A'; c < 'A' + 26; c++ {
			fmt.Printf("%c", c)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()

		for c:='a'; c < 'a' + 26; c++ {
			fmt.Printf("%c",c)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("\nWaiting for goroutines")
	wg.Wait()
	fmt.Println("\nAll goroutines are done!")
}