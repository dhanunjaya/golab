package main

import "fmt"

func max(x int, y int)(int) {
		if x > y {
				return x
		}
		return y
}

func main() {
		x := 3
		y := 4
		z := 10
		max_xy := max(x, y)
		max_xz := max(x, z)
		fmt.Printf("max of %d, %d, %d is %d \n", x, y, z, max(max_xy, max_xz))
}
