package main

import "fmt"

func Max(slice []int) (max int) {
	max = slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return
}

func main() {
	A1 := [7]int {1, 2, 3, 4, 5, 6, 7}
	A2 := [4]int {10, 20, 30, 40}
	A3 := [2]int {70, 5}

	var slice []int
	slice = A1[:]
	fmt.Println("The biggest number is ", Max(slice))
	slice = A2[:]
	fmt.Println("The biggest number is ", Max(slice))
	slice = A3[:]
	fmt.Println("The biggest number is ", Max(slice))
}
