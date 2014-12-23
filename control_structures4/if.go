// This program pritns even or odd for the first 10 numbers

package main

import "fmt"

func main() {
		for i := 1; i <= 10; i++ {
				if i % 2 == 0 {
						fmt.Println(i, "even")
				} else {
						fmt.Println(i, "odd")
				}
		}
}
