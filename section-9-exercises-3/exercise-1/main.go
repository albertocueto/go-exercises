package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Print(i)
		if i == 10000 {
			fmt.Println()
		} else {
			fmt.Print(", ")
		}
	}
}
