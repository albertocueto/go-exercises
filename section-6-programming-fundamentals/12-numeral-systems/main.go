package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("Decimal: %d (%T)\n", n, n)
	fmt.Printf("Binary: %b (%T)\n", n, n)
	fmt.Printf("Hexadecimal: %#X (%T)\n", n, n)
}
