package dotdotdot

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

func PrintAll(kargs ...string){
	fmt.Println("Type of item is:",reflect.TypeOf(kargs))
	for _,value := range(kargs) {
		fmt.Println(value)
	}
	fmt.Println("Overall: ",strings.Join(kargs," "))
}

//this is wired, since we are trying to slow down the goroutines
func RunAll(kargs ...func(int)) {
	fmt.Println("Start running functions...")

	wg := sync.WaitGroup{}
	wg.Add(1)

	for i,f := range kargs{
		go func(idx int){
			defer wg.Done()
			f(idx)
		}(i)
	}
	wg.Wait()
	fmt.Println("Completed running functions")
}
