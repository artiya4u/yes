package main

import (
	"os"
	"bufio"
)

func main() {
	var b = []byte("y\n")
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for {
		f.Write(b)
	}
}
