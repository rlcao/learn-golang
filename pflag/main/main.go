package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)
func main()  {
	var ip *int = flag.Int("flagname", 1234, "help message for flagname")
	flag.Parse()
	fmt.Println("The specified option value is:", *ip)
}
