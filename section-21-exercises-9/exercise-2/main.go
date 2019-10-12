package main

import "fmt"

type person struct {
	name string
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("This is %v speaking\n", p.name)
}
func saySomething(h human) {
	fmt.Printf("The human %v is of type %T\n", h, h)
	h.speak()
}
func main() {
	cueto := person{
		name: "Cueto",
	}
	saySomething(&cueto)
	// saySomething(cueto) // Doesn't work because the receiver is a pointer, person
}

// Receivers       Values
// -----------------------------------------------
// (t T)           T and *T
// (t *T)          *T
