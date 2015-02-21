package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	x := 2
	y := 3
	z := 4

	fmt.Printf("Max of three numbers %d, %d, %d is %d\n", x, y, z, max(max(x, y), z))
}
