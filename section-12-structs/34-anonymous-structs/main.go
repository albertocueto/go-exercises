package main

import "fmt"

// Can be defined here, but we can make an anonymous struct moving it down there:
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

func main() {
	// Composite literal, the best way to refer to this
	// An anonymous struct, it doesn't really have a name
	// Use this when you only need a struct in a very small section.
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Dalinar",
		last:  "Kolin",
		age:   45,
	}
	fmt.Println(p1)
}
