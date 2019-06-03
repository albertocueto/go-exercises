package main

import "fmt"

var a int

// We can create our own ridiculous language:
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// We can't do this:
	// a = b
	// Because they are two different types (int and hotdog). Static language, you can't change types.
	// But we can use conversion:
	a = int(b) // This is not casting, it's called converting. There is no talk about casting in GO, just convert.
	fmt.Printf("%T: %v\n", a, a)
}
