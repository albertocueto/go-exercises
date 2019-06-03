package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Expanding xi slice to fit in the variadic parameter:
	fmt.Println(foo(xi...))
	// This is valid for variadic:
	fmt.Println(foo())
}

// The variadic parameter must be the last one if you were specifying multiple paramters
func foo(x ...int) int {
	fmt.Println("Code from foo:")
	fmt.Println(x)
	fmt.Printf("%T\n", x) // type was a slice of int.

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
