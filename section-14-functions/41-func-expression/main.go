package main

import "fmt"

func main() {
	// In GO, functions are first class citizens, and that means that func is also a type
	f := func() {
		fmt.Println("This is the very first func expression I've written, yey!")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)
}
