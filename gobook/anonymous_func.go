package main

import "fmt"

func main() {
	add := func(x int) int {
		return x + 1
	}
	n := 2
	fmt.Println("After adding 1 to 2 its ", add(n))
}
