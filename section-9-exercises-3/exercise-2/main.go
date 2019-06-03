package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Printf("Number %d (%s):\n", i, string(i))
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\t:", i)
		}
		fmt.Println()
	}
}
