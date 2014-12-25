package main

import "fmt"

func main() {
		var ratings map[string]float32
		ratings = make(map[string]float32)
		ratings["C"] = 5
		ratings["Go"] = 4.5
		ratings["Python"] = 4.5
		ratings["C++"] = 3
		csharp_value, ok := ratings["C#"]
		if ok {
				fmt.Println("C# is in the map and its rating is ", csharp_value)
		} else {
				fmt.Println("C# is not in the map")
		}
}
