package main

import "fmt"

func main() {
	ii := []int{1, 12, 23, 34, 45, 56, 67, 78, 89, 90}
	s := sum(ii...)
	fmt.Println("With numbers:", ii)
	fmt.Println("Normal sum:", s)
	es := evenSum(sum, ii...)
	fmt.Println("Only evens sum:", es)
	os := oddSum(sum, ii...)
	fmt.Println("Only odds sum:", os)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func evenSum(f func(xi ...int) int, vi ...int) int {
	y := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	total := f(y...)
	fmt.Println("Even numbers:", y)
	return total
}

func oddSum(f func(xi ...int) int, vi ...int) int {
	y := []int{}
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	fmt.Println("Odd numbers:", y)
	return f(y...)
}
