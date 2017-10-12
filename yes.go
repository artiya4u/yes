package main

import (
	"os"
)

func main() {
	var b = []byte("y\n")
	for {
		os.Stdout.Write(b)
	}
}
