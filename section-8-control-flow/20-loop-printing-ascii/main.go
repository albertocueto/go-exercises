package main

import "fmt"

func main() {
	start := 33
	end := 122
	for x := start; x <= end; x++ {
		fmt.Printf("%v\t%x\t%#U\t%v\n", x, x, x, string(x))
	}
}
