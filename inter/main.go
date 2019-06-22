package main

import (
	"fmt"
	"strings"
	"time"
)

type  Shape interface {
	Area() float32
}



type  Square struct{
	length float32
	height float32
}

func (this Square) Area() float32 {
	area := this.height * this.length
	fmt.Printf("Calculating Area for Square: %f", area)
	return area
}

func main() {
	var test Shape = Square{1.0,2.0}

	fmt.Printf("Returned Value is: %f",test.Area())



	s := []string{"one","two","three"}
	for _,v := range s {
		go func(a string) {
			fmt.Println(a)
		}(v)
	}
	time.Sleep(time.Second)


	str := "hello"
	str = strings.Replace(str,"h","x",1)
	fmt.Println(str)
}
