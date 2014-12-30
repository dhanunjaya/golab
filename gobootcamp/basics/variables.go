package main

import "fmt"

var (
		first_name string
		age int
		last_name = "naidu"
)

func main() {
		x := "golang"
		fun := func() {
				fmt.Println(x)
		}
		fun()
		first_name = "dhanunjaya"
		age = 29
		fmt.Println("first name ", first_name)
		fmt.Println("last name ", last_name)
		fmt.Println("Age ", age)
}
