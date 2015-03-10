package main

import (
	"fmt"
	"time"
)

func callMe() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d I am in callMe()\n", i)
	}
}

func main() {
	go callMe()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d I am in main()\n", i)
	}
	time.Sleep(5)
}
