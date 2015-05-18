package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("My favourite number is %d\n", rand.Intn(10))
	//fmt.Printf("random number is %d\n", rand.Seed())
}
