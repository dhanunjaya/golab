package main

import "fmt"

func main() {
		var array [2][3]string
		for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
						array[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
				}
		}
		fmt.Printf("%q\n", array)
}
