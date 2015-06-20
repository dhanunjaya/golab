package main

import "fmt"

type Person struct {
	name    string
	company string
}

func main() {
	person := Person{"Dhanunjaya", "AT&T"}
	fmt.Printf("%s is working at %s\n", person.name, person.company)
}
