package main

import "fmt"

type Shaper interface {
		Area() int
}

type Rectangle struct {
		Length int
		Breadth int
}

type Square struct {
		Side int
}

func (R Rectangle) Area() int {
		return R.Length * R.Breadth
}

func (S Square) Area() int {
		return S.Side * S.Side
}

func main() {
		rectangle := Rectangle{Length:2, Breadth:3}
		square := Square{Side:4}
		var s Shaper
		s = rectangle
		fmt.Println("Area of Rectangle ", s.Area())
		s = square
		fmt.Println("Area of Square ", s.Area())
}
