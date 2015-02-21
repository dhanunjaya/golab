package main

import  "fmt"

func add(x int) {
	x = x + 1
}

func main() {
	a := 3
	fmt.Println("value of a befor calling add() ", a)
	fmt.Println("Value of a after calling add() ", add(a))
}
