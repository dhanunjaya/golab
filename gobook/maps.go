package main

import "fmt"

func main() {
	ratings := map[string]int {"C++":1, "Go":4, "C":5, "Python":3}
	var person map[string]int
	person = make(map[string]int)
	person["pradeep"] = 30
	person["Dileep"] = 31
	for key, value := range ratings {
		fmt.Println(key, value)
	}
	//person["Dileep"] = 1, false
	for _, value := range person {
		fmt.Println(value)
	}
	dileep, ok := person["Dileep"]
	if ok {
		fmt.Println("Dileep age is ", dileep)
	} else {
		fmt.Println("I dont know Dileep")
	}
}
