package main

import (
	"github.com/dhanunjaya/arithmetic"
	"fmt"
)

func main() {
	fmt.Println("Add 2 + 3 = ", arithmetic.Add(2, 3))
	fmt.Println("Sub 3 - 2 = ", arithmetic.Sub(2, 3))
	fmt.Println("Mul 2 * 3 = ", arithmetic.Mul(2, 3))
	fmt.Println("Div 3 / 2 = ", arithmetic.Div(2, 3))
}
