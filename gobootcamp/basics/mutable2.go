package main

import "fmt"

type Artist struct {
		Name, Genre string
		Songs int
}

func newRelease(a *Artist) int {
		a.Songs++
		return a.Songs
}

func main() {
		me := &Artist{Name:"Mat", Genre:"Electro", Songs:42}
		fmt.Printf("%s released their %d song\n", me.Name, newRelease(me))
		fmt.Printf("%s total songs are %d\n", me.Name, me.Songs)
}
