package main

import "fmt"

func main() {
	func() {
		fmt.Println("This function's name will be forever lost to history")
	}()

	f := aFunc()
	f()

	aFuncCallingFunc(f)
	aFuncCallingFunc(func() {
		fmt.Println("This function will also be lost in time")
	})
	aFuncCallingFunc(aFunc())
	fmt.Println(enclosureDemoFunc())
}

func aFunc() func() {
	return func() {
		fmt.Println("This function was born from another function")
	}
}

func aFuncCallingFunc(f func()) {
	fmt.Println("Calling function from the function calling function")
	f()
}

func enclosureDemoFunc() string {
	s := "From the ancient times comes an important message from someone that won't be remembered"
	{
		y := "unite them"
		s = fmt.Sprintf("%v: %v", s, y)
	}
	return s
}
