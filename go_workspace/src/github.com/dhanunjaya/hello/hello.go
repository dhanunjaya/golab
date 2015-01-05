package main

import (
	"fmt"
	"github.com/dhanunjaya/arithmetic"
)

var (
	a int
	b int
)

func main() {
	fmt.Println("Enter first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter second number:")
	fmt.Scanf("%d", &b)
	fmt.Printf("Adding %d and %d is %d\n", a, b, arithmetic.Add(a, b))
}
