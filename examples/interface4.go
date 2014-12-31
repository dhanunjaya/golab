package main

import "fmt"

type Animal interface {
		Speak() string
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) Speak() string {
		return "woof!"
}

func (cat Cat) Speak() string {
		return "Meow!"
}

func main() {
		animals := []Animal {Dog{}, Cat{}}
		for _, animal := range animals {
				fmt.Println(animal.Speak())
		}
}
