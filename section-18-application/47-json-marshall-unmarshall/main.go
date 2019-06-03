package main

import (
	"encoding/json"
	"fmt"
)

type author struct {
	Name  string   `json:"Name"`
	Books []string `json:"Books"`
	Alive bool     `json:"Alive"`
}

func main() {
	a1 := author{
		Name:  "Brandon Sanderson",
		Books: []string{"The way of kings", "Radiant words", "Oathbringer", "Mistborn"},
	}

	a2 := author{
		Name:  "Stephen King",
		Books: []string{"It", "Cujo", "The Stand", "The Talisman"},
	}

	authors := []author{
		a1,
		a2,
	}

	fmt.Println("Let's us MARSHALL...")
	fmt.Println(authors)

	bs, err := json.Marshal(authors)
	if err != nil {
		fmt.Println("There was an error marshalling that", err)
	}
	fmt.Printf("%v (%T)\n", bs, bs)
	fmt.Printf("%v\n", string(bs))

	fmt.Println()
	fmt.Println("Let us UNMARSHALL...")
	jsonStringAuthors := `[{"Name":"Brandon Sanderson","Books":["The way of kings","Radiant words","Oathbringer","Mistborn"],"Alive":false},{"Name":"Stephen King","Books":["It","Cujo","The Stand","The Talisman"],"Alive":false}]`
	jsonBSAuthors := []byte(jsonStringAuthors)
	fmt.Printf("%v (%T)\n", jsonStringAuthors, jsonStringAuthors)
	fmt.Printf("%v (%T)\n", jsonBSAuthors, jsonBSAuthors)

	// authors := []author{}
	fmt.Println()
	var unmarshalledAuthors []author
	err = json.Unmarshal(jsonBSAuthors, &unmarshalledAuthors)
	if err != nil {
		fmt.Println("There was an error unmarshalling that", err)
	}
	fmt.Println("Look at all of these authors:")
	for i, v := range authors {
		fmt.Printf("%v: %v\nbooks:\n", i+1, (v).Name)
		for _, book := range (v).Books {
			fmt.Println(book)
		}
		fmt.Println()
	}

}
