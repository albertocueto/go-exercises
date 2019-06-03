package main

import "fmt"

func main() {
	if true {
		fmt.Println("Truth")
	}

	if !true {
		fmt.Println("Truth")
	}

	if false {
		fmt.Println("Lie")
	}

	if !false {
		fmt.Println("Not a Lie")
	}

	if 2 == 2 {
		fmt.Println("2 is 2")
	}

	if 2 != 2 {
		fmt.Println("2 isn't 2, it's a different two")
	}

	if !(2 != 2) {
		fmt.Println("2 is not different to 2")
	}

	if x := 42; x == 2 {
		fmt.Println("42 was initialized in if statement, and it's equal to 2")
	}

	if y := 42; y == 42 {
		fmt.Println("42 was initialized in if statement, and it's equal to 42")
	}
}
