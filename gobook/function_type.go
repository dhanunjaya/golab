package main

import "fmt"

type test_int func(int) bool

func IsOdd(num int) bool {
	if num % 2 == 0 {
		return false
	}
	return true
}

func IsEven(num int) bool {
	if num % 2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f test_int) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("The even list ", filter(slice, IsEven))
	fmt.Println("The odd list ", filter(slice, IsOdd))
}
