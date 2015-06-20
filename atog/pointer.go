package main

import "fmt"

func main() {
	var p *int
	var i = 10
	p = &i
	fmt.Println(*p)
}
