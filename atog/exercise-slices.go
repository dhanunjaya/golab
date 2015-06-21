package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
        mypic := make([][]uint8, dy)
        for i := 0; i < dy; i++ {
                mypic[i] = make([]uint8, dx)
        }
        for x := 0; x < dy; x++ {
                for y := 0; y < dx; y++ {
                        mypic[x][y] = (uint8(x) + uint8(y))/2
                }
        }
        return mypic
}

func main() {
        pic.Show(Pic)
}
