package main

import "fmt"

func PrintIntSlice(name string, slice []int){
	fmt.Printf("%s == [", name)
	for index := 0; index < len(slice)-1; index++ {
		fmt.Printf("%d ", slice[index])
	}
	fmt.Printf("%d]\n", slice[len(slice)-1])
}

func GrowIntSlice(slice []int, add int) []int {
	new_capacity := cap(slice) + add
	new_slice := make([]int, len(slice), new_capacity)
	for index := 0; index < len(slice); index++ {
		new_slice[index] = slice[index]
	}
	return new_slice
}

func main() {
	slice := []int {0, 1, 2, 3}
	PrintIntSlice("slice", slice)
	fmt.Println("Before GrowIntSlice")
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("cap(slice) = ", cap(slice))
	
	slice = GrowIntSlice(slice, 3)
	fmt.Println("After GrowIntSlice")
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("cap(slice) = ", cap(slice))

	slice = slice[:len(slice)+2]
	slice[4] = 4
	slice[5] = 5
	PrintIntSlice("slice", slice)
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("cap(slice) = ", cap(slice))
	
}
