package main

import (
	"bytes"
	"fmt"
	"sort"
)

type author struct {
	Name      string   `json:"Name"`
	Books     []string `json:"Books"`
	Alive     bool     `json:"Alive"`
	BirthYear int      `json:"BirthYear"`
}

func (a author) String() string {
	var buff bytes.Buffer
	isAlive := ""
	if !a.Alive {
		isAlive = "deceased"
	}
	buff.WriteString(fmt.Sprintf("author: %v (%v %v)\nbooks:\n", a.Name, a.BirthYear, isAlive))
	for _, book := range a.Books {
		buff.WriteString(fmt.Sprintf("\t%v\n", book))
	}
	return buff.String()
}

// ByBirthYear Helps sort a slice of authors by BirthYear
type ByBirthYear []author

// Len returns the length of a slice of authors
func (a ByBirthYear) Len() int { return len(a) }

// Swap excahnges two elements in a slice of authors
func (a ByBirthYear) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less tells if an author is less taking into account the BirthYear field
func (a ByBirthYear) Less(i, j int) bool { return a[i].BirthYear < a[j].BirthYear }

// ByName Helps sort a slice of authors by Name
type ByName []author

// Len returns the length of a slice of authors
func (a ByName) Len() int { return len(a) }

// Swap excahnges two elements in a slice of authors
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less tells if an author is less taking into account the Name field
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	xi := []int{4, 7, 3, 24, 67, 2, 1, 34, 67}
	xs := []string{"King", "Howey", "Sanderson", "Abercrombie", "Asimov", "Verne", "Clark", "Gaiman", "Posteguillo", "Navarro", "García", "Saramago"}

	fmt.Println("Unsorted")
	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("Sorted")
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println()
	fmt.Println("Custom sort")
	a1 := author{
		Name:      "Brandon Sanderson",
		Books:     []string{"The way of kings", "Radiant words", "Oathbringer", "Mistborn"},
		Alive:     true,
		BirthYear: 1975,
	}
	fmt.Println(a1)

	a2 := author{
		Name:      "Stephen King",
		Books:     []string{"It", "Cujo", "The Stand", "The Talisman"},
		Alive:     true,
		BirthYear: 1947,
	}
	fmt.Println(a2)

	a3 := author{
		Name:      "Hugh Howey",
		Books:     []string{"Halfway Home", "Molly Fyde", "Wool"},
		Alive:     true,
		BirthYear: 1975,
	}
	fmt.Println(a3)

	a4 := author{
		Name:      "Robert Jordan",
		Books:     []string{"The eye of the world", "A Memory of light", "The Dragon reborn"},
		Alive:     false,
		BirthYear: 1948,
	}
	fmt.Println(a4)

	a5 := author{
		Name:      "Joe Abercrombie",
		Books:     []string{"The blade itself", "Before they are hanged", "Last argument of kings", "Best Served cold"},
		Alive:     false,
		BirthYear: 1974,
	}
	fmt.Println(a5)

	a6 := author{
		Name:      "Carl Zimmer",
		Books:     []string{"Parasite Rex", "Brain cuttings", "She has her mother's laugh", "A planet of viruses"},
		Alive:     false,
		BirthYear: 1966,
	}
	fmt.Println(a6)

	authors := []author{
		a1,
		a2,
		a3,
		a4,
		a5,
		a6,
	}

	fmt.Println("Unsorted")
	fmt.Println(authors)
	fmt.Println()

	fmt.Println("Sorted by year of birth:")
	sort.Sort(ByBirthYear(authors))
	fmt.Println(authors)
	fmt.Println()

	fmt.Println("Sorted by Name:")
	sort.Sort(ByName(authors))
	fmt.Println(authors)

}
