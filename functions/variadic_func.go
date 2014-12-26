package main

import "fmt"

type person struct {
		name string
		age int
}

func Older(slice ...person) (bool, person) {
		if len(slice) == 0 {
				return false, person{}
		}
		max_older := slice[0]
		for _, value := range slice {
				if value.age > max_older.age {
						max_older = value
				}
		}
		return true, max_older
}

func main() {
		var (
				ok bool
				older person
		)

		jim := person{"jim", 30}
		paul := person{"paul", 32}
		john := person{"john", 28}
		kim := person{"kim", 35}

		_, older = Older(jim, paul)
		fmt.Printf("I am the older %s\n", older.name)
		_, older = Older(jim, paul, kim)
		fmt.Printf("I am the older %s\n", older.name)
		_, older = Older(john)
		fmt.Printf("I am %s\n", older.name)
		ok, older = Older()
		if !ok {
				fmt.Println("There is no one to compare")
		}	
}
