package main

import "fmt"

func main() {
	index := 8
	switch index {
	case 4:
		fmt.Println("the integer was 4") 
	case 6:
		fmt.Println("the integer was 6")
	case 8:
		fmt.Println("the integer was 8 and execute fallthrough")
		fallthrough
	case 10:
		fmt.Println("i dont the integer value bcoz of fallthrough i am with you")
		fallthrough
	default:
		fmt.Println("who will care me")
	}
}
