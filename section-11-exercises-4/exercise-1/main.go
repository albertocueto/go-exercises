package main

import "fmt"

func main() {
	// Array, need to specify the size beforehand:
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of our array composite: %T\nThe array: %v\n", x, x)
}
