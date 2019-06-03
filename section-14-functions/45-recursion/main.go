package main

import "fmt"

// Remember, most things you can do with recursion, you can do with loops

func main() {
	fmt.Println("Factorial using recursion")
	fmt.Println(15, factorialRecursion(15))
	fmt.Println(4, factorialRecursion(4))

	fmt.Println()

	fmt.Println("Factorial using a loop")
	fmt.Println(15, factorialLoop(15))
	fmt.Println(4, factorialLoop(4))

	fmt.Println()

	fmt.Println("Factorial using another loop")
	fmt.Println(15, factorialLoopAlternate(15))
	fmt.Println(4, factorialLoopAlternate(4))
}

func factorialRecursion(i int) int {
	if i == 1 {
		return i
	}
	return i * factorialRecursion(i-1)
}

func factorialLoop(i int) int {
	fact := i
	for j := i - 1; j > 1; j-- {
		fact *= j
	}
	return fact
}

func factorialLoopAlternate(n int) int {
	fact := 1
	for ; n > 1; n-- {
		fact *= n
	}
	return fact
}
