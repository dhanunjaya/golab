package main

import "fmt"

func main() {
	array := [...]string{"Hello", "world"}
	fmt.Println(array)
	fmt.Printf("%s\n", array)
	fmt.Printf("%q\n", array)
}
