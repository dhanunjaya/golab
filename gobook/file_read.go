package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Open() %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 100)
	n, err := f.Read(buf)
	stringVersion := string(buf)
	fmt.Printf("%d %s", n, stringVersion)
}
