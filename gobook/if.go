package main

import "fmt"

func compute_x() int {
	return 11
}

func main() {
	if x := compute_x(); x > 10 {
		fmt.Println("x is greater than 10")
	}else {
		fmt.Println("x is less than 10")
	}	
}
