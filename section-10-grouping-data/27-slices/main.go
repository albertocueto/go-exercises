package main

import "fmt"

func main() {
	// x := type{values} <- Composite literal
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Slice. Allows you to group together values of the same type
	fmt.Println(x)
}
