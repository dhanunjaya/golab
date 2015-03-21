package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(ch chan int) {
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-t.C:
			ch = nil
		}
	}
}

func generator(ch chan int) {
	t := time.NewTimer(2 * time.Second)

	for {
		select {
		case ch <- rand.Intn(42):
		case <-t.C:
			ch = nil
		}
	}
}

func main() {
	ch := make(chan int)
	go worker(ch)
	go generator(ch)
	time.Sleep(10 * time.Second)
}
