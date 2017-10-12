package main

import (
	"os"
	"bufio"
)

func main() {
	var b = []byte("y\n")
	f := bufio.NewWriterSize(os.Stdout, 64*1024)
	defer f.Flush()
	for {
		f.Write(b)
	}
}
