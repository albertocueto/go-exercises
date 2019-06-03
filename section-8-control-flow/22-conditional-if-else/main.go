package main

import "fmt"

func main() {
	x := 42
	if x == 42 {
		fmt.Printf("The value %v was equal to 42\n", x)
	} else if x == 41 {
		fmt.Printf("The value %v was actually 40\n", x)
	} else {
		fmt.Printf("The value %v was something completely different\n", x)
	}
}
