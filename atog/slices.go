package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 7}
	fmt.Println("s==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}
}
