package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
	profession
}

type profession struct {
	title        string
	yearsOfStudy int
}

func main() {
	p1 := person{
		first: "Alberto",
		last:  "Cueto",
		age:   33,
		profession: profession{
			title:        "Software Engineer",
			yearsOfStudy: 5,
		},
	}

	p2 := person{
		first: "Diana",
		last:  "Arvizu",
		age:   33,
		profession: profession{
			title:        "MD",
			yearsOfStudy: 5,
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	// Notice that the inner type is promoted to the outer type, so you can either do:
	// p1.title or p1.profession.title
	// In case of collision of names of the attributes, you can specify which you're referring to
	fmt.Println(p1.first, p1.last, p1.age, p1.title, p1.profession.yearsOfStudy)
	fmt.Println(p2.first, p2.last, p2.age, p1.title, p1.profession.yearsOfStudy)
}
