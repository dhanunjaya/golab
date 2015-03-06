package main

import (
	"fmt"
	"errors"
	"os"
)

var (
	errorString = errors.New("unwilling to print empty string\n")
)

func printer(msg string) error {
	if msg == "" {
		return errorString
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer(""); err != nil {
		fmt.Printf("printer error: %s", err)
		os.Exit(1)
	}
}
