package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length float64
	breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{2, 3}
	c1 := Circle{2}
	fmt.Println("Area of a rectangle ", r1.area())
	fmt.Println("Area of a circle ", c1.area())
}
