package main

import "fmt"

func main() {
	var fa float32
	fmt.Println("Enter fahrenheit")
	fmt.Scanf("%f", &fa)
	celsius := (fa - 32) * (5 / 9.0)
	fmt.Println(celsius)	
}
