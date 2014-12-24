package main

import "fmt"

func main() {
		x := [5]uint8{10, 20, 30, 40, 50}
		var total uint8 = 0
		for _, value := range x {
			total += value	
		}
		fmt.Println("Total is ", total)	
}
