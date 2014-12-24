package main

import "fmt"

func main() {
		x := [5]uint8{10, 20, 30, 40, 50}
		var total uint8 = 0
		for i := 0; i < len(x); i++ {
				total += x[i]
		}
		fmt.Println("total is ", total)	
}
