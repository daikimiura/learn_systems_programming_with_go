package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multi-writer.txt")

	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}