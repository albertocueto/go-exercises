package main

import "fmt"

type person struct {
	name string
}

func main() {
	x := 1
	fmt.Println("value:", x, "address:", &x)
	p := person{
		name: "Cueto",
	}
	fmt.Println(p)
	changePerson(&p)
	fmt.Println(p)
}

func changePerson(p *person) {
	// *p = person{
	// 	name: fmt.Sprintf("Super %v", (*p).name),
	// }
	// Dereferencing p:
	(*p).name = fmt.Sprintf("Super %v", (*p).name)
}
