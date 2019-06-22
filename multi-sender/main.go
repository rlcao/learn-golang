package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const NumSenders = 10

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(11)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(index int) {
			defer wgReceivers.Done()
			for {

				select {
				case <- stopCh:
					return
				case dataCh <- index:
				}
			}
		}(i)
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()
        var sum int
		for value := range dataCh {
			sum += value;
			if (sum > 10000) {
				fmt.Println("Total sum is:",sum)
				//stopping
				close(stopCh)
				return
			}
		}
	}()

	// ...
	wgReceivers.Wait()
}
