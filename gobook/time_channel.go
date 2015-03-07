package main

import (
	"fmt"
	"time"
)

func emit(wordCh chan string, doneCh chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
	defer close(wordCh)
	t := time.NewTimer(2 * time.Second)
	i := 0
	for {
		select {
			case wordCh <- words[i]:
				i++
				if i == len(words) {
					i = 0
				}
			case <-doneCh:
				fmt.Printf("I am Done!")
				close(doneCh)
				return
			case <-t.C:
				return
		}
	}	
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for word := range wordCh{
		fmt.Printf("%s ", word)
	}
}
