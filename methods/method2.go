/* Are methods applicable only for struct types? The anwser is No. 
	In fact, you can write methods for any named type that you 
	define, that is not a pointer:
*/

package main

import "fmt"

type SliceOfInts []int
type AgeOfPersons map[string]int

func (s SliceOfInts) sum() int {
		sum := 0
		for _, value := range s {
				sum += value
		}
		return sum
}

func (people AgeOfPersons) older() string {
		v := 0
		k := ""
		for key, value := range people {
				if value > v {
						v = value
						k = key
				}
		}
		return k
}

func main() {
		slice := SliceOfInts {1, 2, 3, 4, 5} 
		folks := AgeOfPersons {"Bob":35, "Mike":42, "John": 28}
		fmt.Println("The sum is ", slice.sum())
		fmt.Println("I am the older of this group ", folks.older())
}
