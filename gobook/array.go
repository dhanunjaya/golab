package main

import "fmt"

type person struct {
	name string
	age int
}

func Older5(people [5]person) (older person) {
	older = people[0]
	for index := 1; index < len(people); index++ {
		if people[index].age > older.age {
			older = people[index]
		}
	}
	return
}

func main() {
	var array [5]person
	array[0] = person{"Bob", 25}
	array[1] = person{"Rob", 35}
	array[2] = person{"Sam", 45}
	array[3] = person{"Jim", 55}
	array[4] = person{"Paul", 15}

	older := Older5(array)
	fmt.Printf("The older one is %s and his age is %d\n", older.name, older.age)
}
