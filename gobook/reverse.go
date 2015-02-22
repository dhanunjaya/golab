package main

import "fmt"

func Reverse(slice []byte) {
	length := len(slice)
	if length > 1 {
		slice[0], slice[length-1] = slice[length-1], slice[0]
		Reverse(slice[1:length-1])
	}
}

func main() {
	slice := []byte {'A', 'M', 'A', 'R'}
	fmt.Printf("slice = %s\n", slice)
	Reverse(slice)
	fmt.Printf("After reverse slice = %s\n", slice)
}
