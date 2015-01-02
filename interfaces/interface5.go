/* An interface type is a set of methods, A value of an interface type
	can hold any value that implements those methods
*/

package main

import "fmt"

type Namer interface {
		Name() string
}

type User struct {
		FirstName string
		LastName string
}

type Customer struct {
		FullName string
		Id int
}

func (u *User) Name() string {
		return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (c *Customer) Name() string {
		return c.FullName
}

func Greet(n Namer) string {
		return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
		u := &User{FirstName: "Dhanunjaya", LastName: "Naidu"}
		c := &Customer{FullName: "Pawan Kalyan"}
		fmt.Println(Greet(u))
		fmt.Println(Greet(c))
}
