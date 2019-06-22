package main

import (
    "dotdotdot"
    "fmt"
    "time"
)

func createFunc() func(int) {
    return func(id int){
        time.Sleep(1*time.Second);
        fmt.Println("My ID is:",id)
    }
}

func main() {
    dotdotdot.PrintAll("This","is","a","test","for","multi","args")
    dotdotdot.RunAll(createFunc(),createFunc(),createFunc(),createFunc(),createFunc())
}
