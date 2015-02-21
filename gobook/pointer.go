package main

import "fmt"

func main() {
	var i = 5
	var ptr *int
	ptr = &i
	name := "dhanunjaya"
	name_ptr := &name
	fmt.Println("value of i ", i)
	fmt.Println("ptr pointing to value ", *ptr)
	fmt.Println("value of name ", name)
	fmt.Println("name_ptr pointing to value ", *name_ptr)
}
