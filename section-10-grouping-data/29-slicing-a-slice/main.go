package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	fmt.Println(x[1])
	fmt.Println(x[1:])  // Slicing from index 1 to the end
	fmt.Println(x[1:5]) // Up to but not including index 5

	fmt.Println("Using for with range to go through the slice")
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("Using for without range to go through the slice")
	for i := 0; i <= 9; i++ {
		fmt.Println(i, x[i])
	}
}
