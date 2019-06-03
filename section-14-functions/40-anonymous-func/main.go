package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Anonymous func also ran")
	}()
	func(x int) {
		fmt.Println("The answer to all of the questions:", x)
	}(42)
}

// Named function:
func foo() {
	fmt.Println("Foo ran")
}
