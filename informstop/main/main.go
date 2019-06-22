package main

import (
	"fmt"
	"time"
)

func main()  {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Got signal from stop channel")
				return
			default:
				fmt.Println("Continue working here")
				time.Sleep(1*time.Second)
			}
		}
	}()


	time.Sleep(10*time.Second)
	fmt.Println("Timeout, stopping")
	stop <- true
	time.Sleep(2*time.Second)
}
