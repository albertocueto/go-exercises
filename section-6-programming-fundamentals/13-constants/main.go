package main

import "fmt"

// const a = 42
// const b = 42.78
// const c = "Kaladin Stormlight"

// Typed constants:
const (
	a int     = 42
	b float64 = 42.78
	c string  = "Kaladin Stormlight"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)
	// a = 52 Cannot be done, a is a constant
}
