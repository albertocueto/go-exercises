package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("\n%v\n", i)
		for j := i; j < 10; j++ {
			fmt.Printf("%v", j)
			if j < 9 {
				fmt.Print(", ")
			} else {
				fmt.Println()
			}
		}
	}
}
