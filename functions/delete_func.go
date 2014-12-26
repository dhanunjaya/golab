// Using the append function to delete an element

package main

import "fmt"

func delete(del int, slice []int) ([]int) {
		switch del {
				case 0: slice = slice[1:]
				case len(slice)-1 : slice = slice[:len(slice)-1]
				default: slice = append(slice[:del], slice[del+1:]...)
		}
		return slice
}

func main() {
		slice := []int {1, 2, 3, 4, 5, 6}
		fmt.Println("Initially the slice is ", slice)
		slice = delete(0, slice)
		fmt.Println("After deleteing the first element in slice ", slice)
		slice = delete(len(slice)-1, slice)
		fmt.Println("After deleting the last element in slice ", slice)
		slice = delete(3, slice)
		fmt.Println("After deleting the 3 element from slice ", slice)
}
