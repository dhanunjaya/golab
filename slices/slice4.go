// how slices grow

package main

import "fmt"

func GrowSlice(slice []byte, add int) ([]byte) {
		capacity := cap(slice)
		new_slice := make([]byte, len(slice), add+capacity)
		copy(new_slice, slice)
		return new_slice
}

func main() {
		slice := []byte {1, 2, 3, 4, 5}
		fmt.Println("slice is ", slice)
		fmt.Println("len is ", len(slice))
		fmt.Println("cap is ", cap(slice))
		slice = GrowSlice(slice, 3)
		fmt.Println("After slice is grown ", slice)
	//	#slice[5] = 6
		fmt.Println("len is ", len(slice))
		fmt.Println("cap is ", cap(slice))
}
