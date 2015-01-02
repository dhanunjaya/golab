/* An interface is defined by a set of methods. A value of interface type
can hold any value that implements those methods
*/

package main

import "fmt"

type Namer interface {
	Name() string
}

type User struct {
	FirstName string
	LastName  string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{FirstName: "Dhanunjaya", LastName: "Naidu"}
	fmt.Println(Greet(u))
}
