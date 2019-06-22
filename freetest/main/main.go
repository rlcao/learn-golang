package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)
import "runtime/pprof"

var cpuprofile = flag.String("cpuprofile","","write cpu profile to file")
var p *int

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f,err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	done := false

	go func(){
		done = true
	}()

	for !done {
	}
	fmt.Println("done!")


	fmt.Printf("%d",p)
}