package main

import (
	"fmt"
	"time"
)

func main() {
	mychan := make(chan bool,2)

	//close(mychan)

	/*	for {
		_,ok := <- mychan
		if ok {
			fmt.Println("Read from chan OK")
		} else {
			fmt.Println("Read from chan NotOK")
		}
		time.Sleep(1*time.Second)
	}*/
	// read multiple time on closed channel
	/*for {
		select {
		case val := <-mychan:
			fmt.Println("got from chan:",val)
		default:
			fmt.Println("Nothing from chan")
		}
		time.Sleep(2 * time.Second)
	}*/

	//range and close on channel
	go func(ch chan bool) {
		for c := range(ch){
			fmt.Println("here is what I got:",c)
		}
	}(mychan)

	for i:=0; i< 3;i++ {
		mychan <- true
	}

	close(mychan)

	go func(ch chan bool) {
		for c := range(ch){
			fmt.Println("here is what I got again:",c)
		}
	}(mychan)
	time.Sleep(1*time.Second)
}
