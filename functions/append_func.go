// func append(slice []T, elements...T) []T

package main

import "fmt"

func main() {
		slice := []int {1, 2, 3}
		fmt.Println("Initially the slice is ", slice)
		fmt.Println("len and capacity is ", len(slice), cap(slice))
		slice = append(slice, 4, 5)
		fmt.Println("After adding 2 element to slice now its ", slice)
		fmt.Println("len and capacity is ", len(slice), cap(slice))
}
