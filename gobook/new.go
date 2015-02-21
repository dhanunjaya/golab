package main

import "fmt"

func main() {
	var sum = 0
	var double_sum *int
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	double_sum = new(int)
	*double_sum = 2 * sum
	fmt.Println("the sum is ", sum)
	fmt.Println("the double of sum is ", *double_sum)
}
