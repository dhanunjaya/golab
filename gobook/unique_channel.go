package main

import "fmt"

func uniqNum(c chan int) {
	for i:=1; i<=10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	uniqChannel := make(chan int)
	go uniqNum(uniqChannel)
	for uniq := range uniqChannel {
		fmt.Println(uniq)
	}
}
