package main

import "fmt"

func main() {
	var slice []int
	if slice == nil {
		fmt.Println("slice == nil")
	}
	fmt.Println("len(slice) == ", len(slice))	
	fmt.Println("cap(slice) == ", cap(slice))
	slice = make([]int, 4)
	fmt.Println("slice == ", slice)
	fmt.Println("len(slice) == ", len(slice))	
	fmt.Println("cap(slice) == ", cap(slice))
	slice[1] = 3
	slice[3] = 2
	fmt.Println("slice == ", slice)

}
