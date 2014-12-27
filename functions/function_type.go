package main

import "fmt"

type test_int func(int)(bool)

func isOdd(num int) bool {
		if num % 2 == 0 {
				return false
		}
		return true
}

func isEven(num int) bool {
		if num % 2 == 0 {
				return true
		}
		return false
}

func filter_slice(slice []int, f test_int) []int {
		var odd_slice []int
		for _, value := range slice {
				if f(value) {
						odd_slice = append(odd_slice, value)
				}
		}
		return odd_slice
}

func main() {
		slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println("At first the slice is ", slice)
		odd_slice := filter_slice(slice, isOdd)
		fmt.Println("Filtered odd slice ", odd_slice)
		even_slice := filter_slice(slice, isEven)
		fmt.Println("Filtered even slice ", even_slice)
}
