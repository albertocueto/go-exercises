package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5))
	fmt.Println(bar([]int{1, 2, 3, 4, 5}))
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
