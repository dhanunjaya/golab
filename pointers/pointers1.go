package main

import "fmt"

var (
	i int = 60000
	name string = "dhanunjaya"
)

func main() {
		i_ptr := &i
		name_ptr := &name
		fmt.Println("My name is ", *name_ptr)
		fmt.Println("My salary is ", *i_ptr)
}
