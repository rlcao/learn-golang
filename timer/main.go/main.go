package main

import (
	"fmt"
	"time"
)

func main() {
	/*var t <- chan time.Time
	t = time.After(5 * time.Second)
	fmt.Println(time.Now())
	go func() {
		fmt.Println(<-t)
	}()

	for i:=0; i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("count:", i)
	}
	fmt.Println("completed counting")
	*/
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second)
	// This goroutine will hang if timer is stopped
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 is stopped")
	}

}


