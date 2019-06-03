package main

import "fmt"

var y = 42

// Declared this variable to be a String, its type can't be redeclared
// This is redundand and compiler recommends not adding the string part, it's inferred from the string that is being assigned
var z string = "Shaken, not stirred"

// GO is a static language, not Dynamic, you can't redeclare variable types. A variable is declared to store values of a certain type only
var a = `Cueto said: "
shaken, 
not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// go is s a static programming language, can't redeclare type:
	// z = 43
	// fmt.Println(z)
	// fmt.Println("%T\n", z)
}
