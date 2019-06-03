package main

import "fmt"

func main() {
	x := 0
	// For statement with a single condition:
	for x < 10 {
		fmt.Print(x)
		x++
	}
	fmt.Println()

	x = 0
	for {
		if x > 9 {
			break
		}
		fmt.Print(x)
		x++
	}
	fmt.Println()
	fmt.Println("Donezzo.")
}
