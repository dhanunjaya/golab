package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	wordmap := make(map[string]int)
	for _, value := range arr {
		wordmap[value]++
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
