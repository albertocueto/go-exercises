package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // Returns the address in memory in which the value is stored.
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // Shows the type "A pointer to an int", specified by *int
	var b = &a
	c := &a
	d := *b // De referencing this addres to get the value
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T: %v at (%v)\n", d, d, &d)
	fmt.Println(*&d) // & gives you the addres and * gives you the value stored at the address
	*b = 43

	// a and b are both storing the same value, they are references to the same address
	fmt.Println(*b)
	fmt.Println(a)

	fmt.Println()
	fmt.Println()
	// Pointers are good when you have a big ammount of data and you don't want to be moving that around.
	// Everything on GO is passed by value
	x := 0
	foo(x)
	fmt.Println(x)

	fmt.Println()
	fmt.Println()
	fmt.Println(x)
	// This is still considered pass by value, as every pass in GO
	bar(&x)
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(ia *int) {
	fmt.Println(ia, *ia)
	*ia = 43
	fmt.Println(ia, *ia)
}
