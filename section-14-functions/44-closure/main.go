package main

import "fmt"

// Closure: enclose scope of a variable. Enclosing so that the variable is limited.

// Package level scope:
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	whatever("Outside main")
	fmt.Println(x)

	// Closure:
	{
		y := "I'll never get out"
		fmt.Println(y)
		whatever(y)
	}
	fmt.Println(x)
	a := incrementor()
	b := incrementor()

	fmt.Println()
	fmt.Println("A:")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println()
	fmt.Println("B:")
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println()
}

func whatever(s string) {
	fmt.Println(x, "whatever", s)
	x++
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
