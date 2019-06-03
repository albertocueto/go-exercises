package main

import "fmt"

func main() {
	x := 17
	fmt.Println(x)
	x = 18
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
	fmt.Println("Hello universe, we are taking a GO class")
	foo()
	n, error := fmt.Println("Printing second, checking println returns")
	fmt.Println(n, error)
	m, _ := fmt.Println("Printing second, checking println returns")
	fmt.Println(m)

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i, "i%2 == 0?", i%2 == 0)
		} else {
			fmt.Println("Odd", i, "i%2 == 0?", i%2 == 0)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("And then we exited")
}

// control flows
// sequence
// loop iterative
// conditional
