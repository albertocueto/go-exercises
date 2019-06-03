package main

import "fmt"

type stormlight int

var x stormlight

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y = int(x)
	fmt.Printf("%T: %v\n", y, y)
}
