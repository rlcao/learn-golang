package main

import "log"

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Llongfile|log.Ldate|log.Lmicroseconds)
}


func main() {
	log.Println("This is a info?")

	log.Fatalln("A Fatal message")

	log.Panicln("A Panic message")

}
