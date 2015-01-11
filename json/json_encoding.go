package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Caught error: ", err)
	}
	fmt.Println(string(b))
}
