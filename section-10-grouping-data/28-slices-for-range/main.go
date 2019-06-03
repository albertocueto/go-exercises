package main

import "fmt"

func main() {
	// x := type{values} <- Composite literal
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Slice. Allows you to group together values of the same type
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])

	// For with a range from slice. Ease of programming
	for i, v := range x {
		fmt.Println(i, v)
	}
}
