package main

import "fmt"

const (
	WHITE = iota,
	BLACK
	GREEN
	YELLOW
	RED
	BLUE
)

type Color byte

type Box struct {
	width float64
	height float64
	depth float64
	color Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return height * width * depth
}

func (b *Box) setColor(c Color) {
	b.color = c
}

func (boxlist BoxList) BiggerBox() Color {
	volume := 0.00
	color := Color(BLACK)
	for _, value := range boxlist {
		if value.Volume() > volume {
			volume = value.Volume()
			color = value.color
		}	
	}
}

func main() {
	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
	}
	color := BiggerBox(boxes)
}
