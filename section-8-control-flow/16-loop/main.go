package main

import "fmt"

func main() {
	for i := 0; i <= 25; i++ {
		fmt.Print(i)
		if i < 25 {
			fmt.Print(", ")
		} else {
			fmt.Println("\nReady or not! Here I go!!!")
		}
	}
}
