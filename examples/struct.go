package main

import "fmt"

type Rectangle struct {
		Length int
		Breadth int
}

func (R Rectangle) Area() int {
		return R.Length * R.Breadth
}

func main() {
		rectangle := Rectangle{Length:10, Breadth:20}
		fmt.Println("Area of a rectangle ", rectangle.Area())
}
