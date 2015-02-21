package main

import "fmt"

func SumAndProduct(x int, y int) (int, int) {
	return x+y, x*y
}

func main() {
	x := 3
	y := 2

	xMINUSy, xTIMESy := SumAndProduct(x, y)
	fmt.Printf("Sum of %d and %d is %d\n", x, y, xMINUSy)
	fmt.Printf("Product of %d and %d is %d\n", x, y, xTIMESy)
}
