package main

import "fmt"

func main() {
	var a []int
	printslice("a", a)

	a = append(a, 0)
	printslice("a", a)

	a = append(a, 1, 2, 3)
	printslice("a", a)
}

func printslice(s string, x []int) {
	fmt.Printf("%s len = %d, cap = %d, %v\n", s, len(x), cap(x), x)
}
