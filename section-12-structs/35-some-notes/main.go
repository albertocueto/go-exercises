package main

import "fmt"

// We create values of a certain type that are stored in variables.
// And those variables have identifiers:
var x int

// Similarly, we create types that are type struct:
type person struct {
	first string
	last  string
}

// You can do the following, even though is ridiculous code
// Type aliases are a bad idea and bad practice:
type foo int

var y foo

const bar int = 42

// We can also do this, it will determine the type
const rab = 42

func main() {
	p1 := person{
		first: "Ellen",
		last:  "Ripley",
	}
	fmt.Println(p1)
	y = 42
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)

	y = rab
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
}
