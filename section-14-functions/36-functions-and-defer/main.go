package main

import "fmt"

func main() {
	anotherFunction()
	yetAnotherFunction("it figures")
	s1 := andThenAReturningFunction("returnee")
	fmt.Println(s1)
	x, y := multipleReturnValuesFunction("José", "José")
	fmt.Println(x, y)

	foo()
	bar()
	fmt.Println("after normal foobar run")

	// Defer makes this run after the containing code has ran
	defer foo()
	bar()
	fmt.Println("foo call was deferred")
}

// func (r receiver) identifier(parameters) (return(s)) {...}

func anotherFunction() {
	fmt.Println("Printing from another function")
}

// Everything in GO is passed by value
func yetAnotherFunction(s string) {
	fmt.Println("printing from yet another function...", s)
}

func andThenAReturningFunction(s string) string {
	return fmt.Sprintf("Hello from a returning function, you passed the value %v", s)
}

func multipleReturnValuesFunction(f string, l string) (string, bool) {
	b := f == l
	a := fmt.Sprintf("%v %v says hello, you are so %v", f, l, b)
	return a, b
}

func foo() {
	fmt.Println("Typical foo")
}

func bar() {
	fmt.Println("bar")
}
