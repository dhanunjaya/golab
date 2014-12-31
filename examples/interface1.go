package main

import "fmt"

type Shaper interface {
		Area() int
}

type Rectangle struct {
		Length int
		Breadth int
}

func (R Rectangle) Area() int {
		return R.Length * R.Breadth
}

func main() {
		rectangle := Rectangle{Length:2, Breadth:3}	
		fmt.Println("Area of a Rectangle ", rectangle.Area())
		s := Shaper(rectangle)
		fmt.Println("Area of shape ", s.Area())
}
