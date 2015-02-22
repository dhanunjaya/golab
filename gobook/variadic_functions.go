package main

import "fmt"

type person struct {
	name string
	age int
}	

func Older(people ...person)(person, bool) {
	if len(people) == 0 {
		return person{}, false
	}
	older := people[0]
	for _, value := range people {
		if value.age > older.age {
			older = value
		}
	}
	return older, true
}

func main() {
	var (
		older person
		ok bool
	)
	sam := person{"Sam", 23}
	bob := person{"Bob", 34}
	rob := person{"Rob", 42}

	older, _ = Older(sam, bob)
	fmt.Println("The older of sam and bob is : ", older.name)
	older, _ = Older(sam, bob, rob)
	fmt.Println("The older of sam, bob and rob is : ", older.name)
	older, ok = Older()
	if !ok {
		fmt.Println("Hey may be you have sent nothing or couldt find him")
	}
}
