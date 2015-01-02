package main

import "fmt"

type Stringer interface {
		String() string
}

type fakestring struct {
		content string
}

func (s *fakestring) String() string {
		return s.content
}

func printString(value interface{}) {
		switch str := value.(type) {
		case string: fmt.Println(str)
		case Stringer: fmt.Println(str.String())
		}
}

func main() {
		str := &fakestring{content:"sachin ramesh tendulkar"}	
		printString(str)
		printString("virat kohli")
} 
