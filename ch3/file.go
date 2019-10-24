package main

import (
	"io"
	"os"
)

func main() {
	p, _ := os.Getwd()
	file, err := os.Open(p + "/ch3/file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
