package main

import "fmt"

func main() {
		ratings := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}
		// var ratings = map[string]float32
		// ratings = make(map[string]float32)
		for key, value := range ratings {
				fmt.Printf("%s raating is at %f\n", key, value)
		}
}
