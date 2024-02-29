package main

import (
	"os"
	"io"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)
	io.Copy(os.Stdout, file)
}
