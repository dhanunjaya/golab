package main

import "fmt"

func seqNumber(c chan int) {
	numSlice := []int{1, 2, 3, 4, 5}
	for _, num := range numSlice {
		c <- num
	}
	close(c)
}

func main() {
	numChannel := make(chan int)
	go seqNumber(numChannel)
	for value := range numChannel {
		fmt.Println(value)
	}
}
