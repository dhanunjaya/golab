// Example-2
// S is a valid Implementation for I, because it defines the two methods which I requires. 
// Note that this is true even though there is no explicit declaration that S implements I

package main

import "fmt"

type I interface {
	Get() int
	Put(int)
}

type S struct { 
	i int 
}

func (p *S) Get() int { 
	return p.i 
}

func (p *S) Put(v int) { 
	p.i = v 
}

func f(p I) {
	p.Put(1)
	fmt.Println(p.Get())
}

func main() {
	s := S{i:0}
	f(&s)
}
