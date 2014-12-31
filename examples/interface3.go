package main

import "fmt"

type Shaper interface {
		Area() (string, int)
}

type Rectangle struct {
		Length int
		Breadth int
}

type Square struct {
		Side int
}

func (R Rectangle) Area() (string, int) {
		return "Area of Rectangle ", R.Length * R.Breadth
}

func (S Square) Area() (string, int) {
		return "Area of Square ", S.Side * S.Side
}

func main() {
		rectangle := Rectangle{Length:2, Breadth:3}
		square := Square{Side:2}
		slice := []Shaper {rectangle, square}
		for i, _ := range slice {
				str, val := slice[i].Area()
				fmt.Println(str, val)
		}
}
