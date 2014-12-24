/* The idea now is to write a function Older10 that takes an 
	array of 10 persons as its input, and returns the 
	older person in this array.
*/

package main

import "fmt"

type person struct {
		name string
		age int
}

func Older10(array [10]person)(person) {
		older := array[0]
		for index := 1; index < len(array); index++ {
				if array[index].age > older.age {
						older = array[index]
				}	
		}
		return older
}

func main() {
		var array [10]person
		array[1] = person{"sam", 20}
		array[2] = person{"bob", 30}
		array[5] = person{"alice", 25}
		array[6] = person{"john", 28}
		array[7] = person{"david", 32}
		older := Older10(array)
		fmt.Printf("Hey Iam the older guy %s and my age is %d\n", older.name, older.age)
}
