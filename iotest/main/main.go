package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var bf bytes.Buffer;
	bf.Write([]byte("Hello"))

	fmt.Fprint(&bf, " World!")

	io.Copy(os.Stdout, &bf)
}
