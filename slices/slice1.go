package main

import "fmt"

func main() {
		array := [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'k', 'l'}
		var a_slice, b_slice []byte
		a_slice = array[:5]
		b_slice = array[5:len(array)]
		fmt.Println(a_slice)
		fmt.Println(b_slice)
}
