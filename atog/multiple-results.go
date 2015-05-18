package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a := "hello"
	b := "world"
	fmt.Printf("Before swapping a : %s, b : %s\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After swapping a : %s, b : %s\n", a, b)
}
