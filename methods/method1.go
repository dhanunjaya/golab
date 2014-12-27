/* A method is function that is bound or attached to a given type. Its 
	syntax is the same as a traditional function except that we specify 
	a receiver of this type just after the keyword func
*/

package main

import (
		"fmt"
		"math"
)

type Rectangle struct {
		width float64
		height float64
}

type Circle struct {
		radius float64
}

func (r Rectangle) area() (float64) {
		return r.width * r.height
}

func (c Circle) area() (float64) {
		return c.radius * c.radius * math.Pi
}

func main() {
		rectangle := Rectangle{12, 2}
		circle := Circle{2}
		fmt.Println("Area of rectangle is ", rectangle.area())
		fmt.Println("Area of circle is ", circle.area())
}
