package main

import "fmt"

const (
	_ = iota // 0 iota
	// kb = 1024
	kb = 1 << (iota * 10) // 1 iota
	mb = 1 << (iota * 10) // 2 iota
	gb = 1 << (iota * 10) // 3 iota
	tb = 1 << (iota * 10) // 4 iota
)

func main() {
	// golang has the << >> bit shifting parameters:
	x := 2
	fmt.Printf("Decimal: %d\t\tBinary: %b\n", x, x)
	y := x << 1
	fmt.Printf("Decimal: %d\t\tBinary: %b\n", y, y)
	fmt.Printf("Decimal: %d\t\tBinary: %b\n", y<<1, y<<1)
	fmt.Printf("Decimal: %d\t\tBinary: %b\n", y<<1<<1, y<<1<<1)

	kb1 := 1024
	mb1 := 1024 * kb
	gb1 := 1024 * mb

	fmt.Printf("KB Decimal: %d \t\t Binary: %b\n", kb1, kb1)
	fmt.Printf("MB Decimal: %d \t\t Binary: %b\n", mb1, mb1)
	fmt.Printf("GB Decimal: %d \t\t Binary: %b\n", gb1, gb1)

	// Using iota:
	fmt.Println("Using iota:")
	fmt.Printf("KB Decimal: %d \t\t Binary: %b\n", kb, kb)
	fmt.Printf("MB Decimal: %d \t\t Binary: %b\n", mb, mb)
	fmt.Printf("GB Decimal: %d \t\t Binary: %b\n", gb, gb)
	fmt.Printf("TB Decimal: %d \t Binary: %b\n", tb, tb)
}
