package main

import "fmt"

func main() {
	s := "Hi class"
	fmt.Printf("%T: %v\n", s, s)
	// Raw string literal, will include spaces and new lines.
	s = `"Hi class,
	you need to step up your programming skills."`
	fmt.Printf("%T: %v\n", s, s)
	bs := []byte(s)
	fmt.Printf("%T: %v\n", bs, bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

	// Everything you do in GO is utf8 encoded
}
