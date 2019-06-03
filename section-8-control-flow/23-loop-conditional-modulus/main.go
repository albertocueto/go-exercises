package main

import "fmt"

func main() {
	limit := 1000
	for x := 0; x <= limit; x++ {
		if x%2 == 0 {
			fmt.Printf("%v\t", x)
		}
	}
}
