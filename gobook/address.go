package main

import "fmt"

var (
	i int = 9
	name string = "dhanunjaya"
	pi float32 = 3.14
	c complex64 = 3+5i
)

func main() {
	fmt.Println("Hexadecimall address of i ", &i);
	fmt.Println("Hexadecimal address of name ", &name);
	fmt.Println("Heexadecimal address of pi ", &pi);
	fmt.Println("Hexadecimal address of c ", &c);
}
