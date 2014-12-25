// slices are resizable

package main

import "fmt"

func main() {
		var slice []byte
		if slice == nil {
				fmt.Println("slice == nil")
		}
		fmt.Println("slice = ", slice)
		fmt.Println("len is ", len(slice))
		fmt.Println("cap is ", cap(slice))
		slice = make([]byte, 3, 3)
		fmt.Println("slice after make = ", slice)
		fmt.Println("len is ", len(slice))
		fmt.Println("cap is ", cap(slice))
}
