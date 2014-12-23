package main

import "fmt"

func main() {
	var feet float32
	const x = 0.3048
	fmt.Println("Enter feet")
	fmt.Scanf("%f", &feet)
	meter := feet * x
	fmt.Println(feet, " feet is ", meter, " meters")
}
