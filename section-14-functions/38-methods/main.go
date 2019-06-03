package main

import "fmt"

type character struct {
	first string
	last  string
}

type knightRadiant struct {
	character
	order string
}

// Wea re attaching this function to the receiving type:
func (kr knightRadiant) speak() {
	fmt.Printf("Hi, my name is %v %v, I'm a member of the Knight Radiants, from the order of the %v.\nNice to meet you.\n", kr.first, kr.last, kr.order)
}

func main() {
	kr1 := knightRadiant{
		order: "Windrunners",
		character: character{
			first: "Kaladin",
			last:  "Stormblessed",
		},
	}
	fmt.Println(kr1)

	kr2 := knightRadiant{
		order: "Lightweavers",
		character: character{
			first: "Shallan",
			last:  "Davar",
		},
	}
	kr1.speak()
	kr2.speak()
}
