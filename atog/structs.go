package main

import "fmt"

type Person struct {
	name    string
	job     string
	company string
}

func main() {
	fmt.Println(Person{"dhanunjaya", "devops", "AT&T"})
}
