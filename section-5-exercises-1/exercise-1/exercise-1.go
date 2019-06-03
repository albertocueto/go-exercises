package main

import "fmt"

/*
Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
42
“James Bond”
true
Now print the values stored in those variables using
a single print statement
multiple print statements
*/

func main() {
	x := 32
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("%v\t%v\t%v\t\n", x, y, z)
}
