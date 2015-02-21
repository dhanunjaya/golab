package main

import "fmt"

func main() {
	array := [2][4]int {{1, 2, 3, 4}, {5, 6, 7, 8}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("array[%d][%d] = %d\n", i, j, array[i][j])
		}
	}
}
