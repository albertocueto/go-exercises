package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	x = append(x, 10, 11, 12, 13, 14, 15)
	fmt.Println(x)

	y := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y = append(y, 10, 11, 12, 13, 14, 15)

	for _, v := range y {
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	x = append(x[:2], x[7:]...)
	fmt.Println(x)

	y = append(y, x...)
	fmt.Println(y)

	// Slice with a make:
	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[8] = 9
	z = append(z, 10)
	fmt.Println(z)
	fmt.Println("len(z)", len(z))
	fmt.Println("cap(z)", cap(z))
	z = append(z, 11)
	fmt.Println(z)
	fmt.Println("len(z)", len(z))
	fmt.Println("cap(z)", cap(z))
	z = append(z, 12)
	fmt.Println(z)
	fmt.Println("len(z)", len(z))
	fmt.Println("cap(z)", cap(z))
}
