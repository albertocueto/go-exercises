package main

import (
	"fmt"
)

type character struct {
	first string
	last  string
}

type knightRadiant struct {
	character
	order string
}

// keyword identifier type, keyword identifier type, keyword identifier type
// A value can be of more than one type
// With this interface declaration, every type that implements the speak method, becomes also of type human
// Polymorphism!!
type human interface {
	speak()
}

// If you don't add any method to an interface, everything implements it!
// Any type has 0 methods, in addition to other methods, so, the condition to become a type of this interface is met:
type everything interface{}

// Wea re attaching this function to the receiving type:
func (kr knightRadiant) speak() {
	fmt.Printf("Hi, my name is %v %v, I'm a member of the Knight Radiants, from the order of the %v.\nNice to meet you.\n", kr.first, kr.last, kr.order)
}

func (ch character) speak() {
	fmt.Printf("Hi, I'm %v %v, I don't belong to any ancient order of people with special powers, and that's ok.\nNice to meet you.\n", ch.first, ch.last)
}

func reflectOnExistance(e everything) {
	fmt.Printf("I %v, exist!!!!\n", e)
}

// In this function we are using assertion to be able to treat the h as the specific types, inside the switch to give an according response:
func greetHuman(h human) {
	h.speak()
	fmt.Printf("Hi %v (%T), nice to meet you\n", h, h)
	// We are going to use assertion to greet different types of person:
	switch h.(type) {
	case character:
		fmt.Printf("it's so awesome to get to meet youm %v, I've heard great things about you.\n", h.(character).first)
	case knightRadiant:
		fmt.Printf("Wooow, %v, it's so awesome to get to meet someone from the Knight Radiants!! Even more, one of the %v.\n", h.(knightRadiant).first, h.(knightRadiant).order)
	}
}

func main() {
	kr1 := knightRadiant{
		order: "Windrunners",
		character: character{
			first: "Kaladin",
			last:  "Stormblessed",
		},
	}

	kr2 := knightRadiant{
		order: "Lightweavers",
		character: character{
			first: "Shallan",
			last:  "Davar",
		},
	}

	ch1 := character{
		first: "Adolin",
		last:  "Kolin",
	}

	greetHuman(kr1)
	fmt.Println()
	greetHuman(ch1)
	fmt.Println()
	greetHuman(kr2)
	fmt.Println()
	fmt.Println()
	reflectOnExistance(kr1)
	reflectOnExistance(kr2)
	reflectOnExistance(ch1)
}
