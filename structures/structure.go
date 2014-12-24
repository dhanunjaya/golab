/* Let's write our Older function that takes two input parameters of type person,
   and returns the older one and the difference in their ages.
*/

package main

import "fmt"

type person struct {
		name string
		age int
}

func Older(p1 person, p2 person)(person, int) {
		if p1.age > p2.age {
				return p1, p1.age - p2.age
		}
		return p2, p2.age - p1.age		
}

func main() {
		sachin := person{"sachin", 42}
		arjun := person{age:16, name:"arjun"}
		older, age_diff := Older(sachin, arjun)
		fmt.Printf("%s older than %s with age diff of %d\n", older.name,
									arjun.name, age_diff)
}
