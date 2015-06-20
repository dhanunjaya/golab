package main

import "fmt"

type Person struct {
	name    string
	company string
}

func main() {
	person := Person{"Dhanunjaya", "AT&T"}
	personPtr := &person
	personPtr.name = "Pradeep"
	fmt.Printf("%s is working at %s\n", person.name, person.company)
}
