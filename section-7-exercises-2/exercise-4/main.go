package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %#x\n", x, x, x)
	x = x << 1
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %#x\n", x, x, x)
}
