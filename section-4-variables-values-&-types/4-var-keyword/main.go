package main

import "fmt"

// The scope of z is the entire program:
// Declare and assign = initialization
var z = 450

// Declares that there's a variable with a w identifier and that the type is int
// Automatic assignation of ZERO VALUE of type int (0)
// zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var w int

func main() {
	// Declaring a variable and assigning a value
	// Short declaration operator:
	// Can be used between a function body
	// It's a good practice to use these because it narrows the variables' scope:
	x := 42
	fmt.Println(x)
	// The difference is that if you need to declare variables outside function bodies, you do it like this:
	var y = 43
	fmt.Println(y)
	foo()
	fmt.Println("w", w)
}

func foo() {
	fmt.Println(z)
}
