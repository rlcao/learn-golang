package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	//define array test
	array := [...]int{1,2,3,4,5,6}
	multi_dimension_array := [2][2]int{{1,2},{3,4}}

	fmt.Println(array[0])
	fmt.Println("This is a test:", multi_dimension_array[0][1])

	array_string := [...]string{"test","name"}
	fmt.Println("This is another test:", array_string[0])

	//define pointers
	int_pointer_array := [...]*int{new(int),new(int),new(int)}
	fmt.Println("Int array test:", *int_pointer_array[0])

	int_slice := make([]int,5)
	fmt.Println("this is a slice testing:", int_slice[0])


	//for statement
	for_slice := []int{0,1,2,3,4,6}
	for index,value := range(for_slice) {
		fmt.Println("index:",index,"value:",value)
	}


    //hard way to print hello world
    var b bytes.Buffer;

	b.Write([]byte("Hello"))
	fmt.Fprintf(&b, "World!")

	io.Copy(os.Stdout,&b)

	//bslice := []byte("Test") //type converter

	fmt.Fprintf(os.Stdout,"tests")


}