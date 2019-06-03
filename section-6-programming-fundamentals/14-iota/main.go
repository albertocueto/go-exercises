package main

import "fmt"

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

const (
	a = iota
	b
	c
)

const (
	d = iota // Here iota is resetted
	e
	f
)

const (
	a1 = iota
	b1
	c1
	d1
	e1
	f1
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e1)
	fmt.Println(f1)
}
