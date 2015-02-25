package main

import "fmt"

func isOdd(x int) bool {
	if x % 2 == 0 {
		return false
	}
	return true
}

func isBiggerThan5(x int) bool {
	if x >= 5 {
		return true
	}
	return false
}

func filter_list(f func (int) bool) (func (s []int) (yes, no []int)) {
	return func (s []int) (yes, no []int) {
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
	slice := []int {1, 2, 3, 4, 5, 6, 7, 8}
	odd_even_function := filter_list(isOdd)	
	odd, even := odd_even_function(slice)
	fmt.Println("odd list ", odd)
	fmt.Println("even list ", even)
	greater_lesser_function := filter_list(isBiggerThan5)
	greater_5, less_5 := greater_lesser_function(slice)
	fmt.Println("greater than 5 ", greater_5)
	fmt.Println("lesser than 5 ", less_5)
}
