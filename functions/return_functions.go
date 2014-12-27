// Functions that return functions

package main

import "fmt"

func isOdd(num int) bool {
		if num % 2 == 0 {
				return false
		}
		return true
}

func isBiggerThan4(num int) bool {
		if num > 4 {
				return true
		}
		return false
}

func filter_factory(f func(num int) bool) (func (s []int) (yes, no []int)) {
		return func(s []int) (yes, no []int) {
				for _, value := range s {
						if f(value) {
								yes = append(yes, value)
						} else {
								no = append(no, value)
						}
				}
				return
		}
}

func main() {
		slice := []int {1, 2, 3, 4, 5, 6, 7}
		fmt.Println("At first the silce is ", slice)
		odd_even_func := filter_factory(isOdd)
		is_odd, is_even := odd_even_func(slice)
		fmt.Println("The two slices ", is_odd, is_even)
		is_bigger, is_smaller := filter_factory(isBiggerThan4)(slice)
		fmt.Println("The two slices ", is_bigger, is_smaller)
}
