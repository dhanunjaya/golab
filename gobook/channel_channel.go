package main

import (
	"fmt"
	"time"
)

func emit(chanChannel chan chan string, doneCh chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
	wordCh := make(chan string)
	defer close(wordCh)
	chanChannel <- wordCh
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
				close(doneCh)
				return
			case <-t.C:
				return
		}
	}
}

func main() {
	chanCh := make(chan chan string)
	doneCh := make(chan bool)
	go emit(chanCh, doneCh)
	wordCh := <-chanCh
	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
}
