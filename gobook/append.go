package main

import "fmt"

func delete(pos int, slice []int) []int {
	slice = append(slice[:pos], slice[pos+1:]...)
	return slice
}

func main() {
	slice := []int {1, 2, 3, 4, 5, 6, 7}
	//slice = delete(0, slice)
	slice = delete(6, slice)
	slice = delete(4, slice)
	fmt.Println("slice is ", slice)
}
