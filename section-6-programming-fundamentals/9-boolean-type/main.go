package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Printf("A song of %v(%T) and %v(%T)\n", a, a, b, b)
	fmt.Printf("a == b? %v\n", a == b)
	fmt.Printf("a > b? %v\n", a > b)
	fmt.Printf("a < b? %v\n", a < b)
}
