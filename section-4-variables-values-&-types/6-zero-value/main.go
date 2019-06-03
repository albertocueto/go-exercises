package main

import "fmt"

var y string
var x int
var z float32

func main() {
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Jack Sawyer, meddling in the territories"
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
