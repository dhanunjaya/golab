package main

import "fmt"

func emit(wordCh chan string, doneCh chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	for {
		select {
			case wordCh <- words[i]:
				i++
				if i == len(words) {
					i = 0
				}
			case <-doneCh:
				fmt.Println("done\n");
				close(doneCh)
				return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
	}
	doneCh <- true
}
