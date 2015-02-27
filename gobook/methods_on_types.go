package main

import "fmt"

type SliceOfInts []int
type AgesByName map[string]int

func (s SliceOfInts) sum() int {
	sum := 0
	for _, value := range s {
		sum = sum + value
	}
	return sum
}

func (agebynames AgesByName) older() string {
	age := 0
	name := ""
	for key, value := range agebynames {
		if value > age {
			age = value
			name = key
		}
	}
	return name
}

func main() {
	s := SliceOfInts {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	agesbynames := AgesByName {
		"Bob" : 36,
		"Mike" : 44,
		"Jane" : 30,
		"Pop" : 100,
	}
	fmt.Printf("Sum of ints is %d\n", s.sum()) 
	fmt.Printf("I am the older %s\n", agesbynames.older())
}
