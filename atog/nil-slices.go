package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice, len(slice), cap(slice))
	if slice == nil {
		fmt.Println("nil!")
	}
}
