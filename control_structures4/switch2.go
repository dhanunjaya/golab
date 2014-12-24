package main

import "fmt"

func main() {
		var i uint8
		fmt.Println("Enter number 1-5")
		fmt.Scanf("%d", &i)
		switch i {
				case 1 : fmt.Println("One")
						 fallthrough
				case 2 : fmt.Println("Two")
//						 fallthrough
				case 3 : fmt.Println("Three")
						 fallthrough
				case 4 : fmt.Println("Four")
						 fallthrough
				case 5 : fmt.Println("Five")
						 fallthrough
				default : fmt.Println("Unknown number")
		}			
}
