package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favoriteIceCreams []string
}

func main() {
	p1 := person{
		firstName: "José Alberto",
		lastName:  "Cueto Bárcenas",
		favoriteIceCreams: []string{
			"oreo cookies",
			"mint chocolate",
			"wine",
		},
	}
	p2 := person{
		firstName: "Diana Marisol",
		lastName:  "Arvizu Velasco",
		favoriteIceCreams: []string{
			"Ferrero Rocher",
			"mango",
			"lime",
		},
	}

	people := []person{p1, p2}

	peopleMap := map[string]person{}

	for _, pv := range people {
		fmt.Printf("%v's favorite icecreams are:\n", pv.firstName)
		peopleMap[pv.lastName] = pv
		for i, v := range pv.favoriteIceCreams {
			fmt.Printf("Choice %v: %v\n", i+1, v)
		}
		fmt.Println()
	}

	fmt.Println("Again, but printing from the people map:")

	for _, p := range peopleMap {
		fmt.Printf("%v's favorite icecreams are:\n", p.firstName)
		for i, v := range p.favoriteIceCreams {
			fmt.Printf("Choice %v: %v\n", i+1, v)
		}
		fmt.Println()
	}
}
