package main

import "fmt"

const (
	a int = 1
	b     = "B"
)

func main() {
	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
}
