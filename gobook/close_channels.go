package main

import (
	"fmt"
	"time"
)

func printer(msg string, stopCh chan bool) {
	<-stopCh
	fmt.Println(msg)
}

func main() {
	stopCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer %d", i), stopCh)
	}
	//	time.Sleep(5 * time.Second)
	close(stopCh)
	time.Sleep(5 * time.Second)
}
