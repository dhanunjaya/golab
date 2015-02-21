package main

import "fmt"

func main() {
	index := 10
	switch index {
	case 2:
		fmt.Println("the integer was 2")
	case 4:
		fmt.Println("the interger was 4")
	case 6:
		fmt.Println("the integer was 6")
	case 8:
		fmt.Println("the integer was 8")
	case 10:
		fmt.Println("the integer was 10")
	default:
		fmt.Println("all i know is index is integer")
	}
}
