package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println("the sum of first 10 numbers ", sum)
}
