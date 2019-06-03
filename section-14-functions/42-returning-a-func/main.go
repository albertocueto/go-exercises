package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s)

	x := func() int {
		return 451 // Farenheit 451
	}
	fmt.Printf("x has type %T\n", x)
	i := x()
	fmt.Println(i)
	barVar := bar()
	fmt.Printf("barVar has type %T\nreturned value from calling it: %v\n", barVar, barVar())
	fmt.Println(barVar())
	fmt.Println(bar()()) // The func is returned and then executed
}

func foo() string {
	x := "Hello world"
	return x
}

// Function that returns a function that returns an int
func bar() func() int {
	return func() int {
		return 154
	}
}
