package main

import "fmt"

type S struct { 
	i int 
}

func (p *S) Get() int { 
	return p.i 
}

func (p *S) Put(v int) { 
	p.i = v 
}

func main() {
	s := S{i:0}
	s.Put(3)
	fmt.Println(s.Get())
}
