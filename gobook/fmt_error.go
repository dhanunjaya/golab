package main

import (
	"fmt"
	"os"
)

func printer(msg string) error {
	if msg == "" {
		return fmt.Errorf("unwilling to print empty string\n")
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	msg := ""
	if err := printer(msg); err != nil {
		fmt.Printf("printer error : %s", err)
		os.Exit(1)
	}	
}
