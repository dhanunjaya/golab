// Anonymous functions i.e functions without a name

package main

import "fmt"

func main() {
		add1 := func(num int) int {
				return num + 1
		}
		num := 2
		fmt.Println("Anonymous function called ", add1(num))
}
