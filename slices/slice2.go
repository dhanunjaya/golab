/* Say we need to find the biggest value in an array of 10 elements. Good, 
	we know how to do it. Now, imagine that our program has the task of 
	finding the biggest value in many arrays of different sizes! Aha! See? 
	One function is not enough for that, because, remember, 
	the types: [n]int and [m]int are different! They can not be used as 
	an input of a single function.

	Slices fix this quite easily! In fact, we need only write a function 
	that accepts a slice of integers as an input parameter, and 
	create slices of arrays.
*/

package main

import "fmt"

func max(slice []byte) (max byte) {
		max = slice[0]
		for index := 1; index < len(slice); index++ {
				if slice[index] > max {
						max = slice[index]
				}
		}
		return max	
}

func main() {
		A := [6]byte {1, 2, 3, 4, 5, 6}
		B := [4]byte {10, 20, 30, 40}
		C := [9]byte {21, 11, 8, 41, 52, 34, 5, 1, 81}
		var slice []byte
		slice = A[:]
		fmt.Printf("Max value in array A is %d\n", max(slice))
		slice = B[:]
		fmt.Printf("Max value in array B is %d\n", max(slice))
		slice = C[:]
		fmt.Printf("MAx value in array C is %d\n", max(slice))
}
