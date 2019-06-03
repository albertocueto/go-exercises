package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("The result of %v %% 4 = %v\n", x, x%4)
	}
}
