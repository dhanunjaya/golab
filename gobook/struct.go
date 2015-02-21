package main

import "fmt"

type person struct {
	name string
	age int
}

func Older(p1 person, p2 person) (string, int) {
	if p1.age > p2.age {
		return p1.name, p1.age - p2.age
	}
	return p2.name, p2.age - p1.age
}

func main() {
	var tim person
	tim.name = "Tim"
	tim.age = 60

	jim := person{"Jim", 50}
	bob := person{age:30, name:"Bob"}
	
	name, age := Older(jim, bob)
	fmt.Printf("The older one in %s, %s is %s age diff is %d\n", jim.name, bob.name, name, age)
}
