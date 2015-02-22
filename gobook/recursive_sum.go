package main

import "fmt"

func Sum(slice []int) int {
	length := len(slice)
	total := 0
	if length > 0 {
		total = total + slice[0]
		return total + Sum(slice[1:])
	}
	return total
}

func main() {
	slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Sum is ", Sum(slice))
}
