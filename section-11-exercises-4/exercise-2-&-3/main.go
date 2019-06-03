package main

import "fmt"

func main() {
	x := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, v := range x {
		fmt.Printf("Value at position %v: \"%v\"\n", i, v)
	}
	fmt.Printf("Type of the slice %v: \"%T\"\n", x, x)
	slice1 := x[:6]
	slice2 := x[6:10]
	// Slicing does throws a slice index out of bounds
	// slice3 := x[6:23]
	slice4 := x[3:8]
	fmt.Println(slice1)
	fmt.Println(slice2)
	// fmt.Println(slice3)
	fmt.Println(slice4)
}
