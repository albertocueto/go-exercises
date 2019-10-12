package main

import "fmt"

type Pipeline struct {
	In  chan<- *string
	Out <-chan *string
}

type PipelineMap map[string]map[string]*Pipeline

func main() {
	// Maps are unordered lists of key, value pairs
	// COMPOSITE LITERAL
	m := map[string]int{
		"José":     1,
		"Alberto":  2,
		"Cueto":    3,
		"Bárcenas": 4,
	}
	fmt.Println(m)
	fmt.Println("Cueto ", m["Cueto"])
	fmt.Println("Unexistent ", m["Unexistent"])
	v, ok := m["Unexistent"]
	fmt.Println(v, ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Doesn't exist")
	}
	// Shorter:
	if v, ok := m["Cueto"]; ok {
		fmt.Println("Cueto is in the map: ", v)
	}

	// Ranged version of map
	for k, v := range m {
		fmt.Printf("Key: %v, value: %v\n", k, v)
	}

	slicex := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range slicex {
		fmt.Println(i, v)
	}

	// Deleting an entry from a map:
	m["wrong_value"] = 5

	fmt.Println("Before deleting value")
	for k, v := range m {
		fmt.Printf("Key: %v, value: %v\n", k, v)
	}

	delete(m, "wrong_value")
	delete(m, "non_existent_value")

	fmt.Println("After deleting value")
	for k, v := range m {
		fmt.Printf("Key: %v, value: %v\n", k, v)
	}

	fmt.Println(m["non_existent_value"])
	m["non_existent_value"] = 6

	if v, ok := m["non_existent_value"]; ok {
		fmt.Println("About to delete value: ", v)
		delete(m, "non_existent_value")
	}

	fmt.Println(m)

	improvedPipelines := PipelineMap{
		"movies": {
			"insert": &Pipeline{
				In:  make(chan *string, 5),
				Out: make(chan *string, 5),
			},
			// "update": updateMoviePipeline(accountMapper, db),
			// "delete": deleteMoviePipeline(accountMapper, db)
		},
		"movies_provider_labels": {
			"update": &Pipeline{
				In:  make(chan *string, 5),
				Out: make(chan *string, 5),
			},
			// "delete": deleteMovieProviderLabelPipeline(accountMapper, db),
		},
	}

	for k, v := range improvedPipelines {
		fmt.Printf("%v\n, %v\n", k, v)
		for l, m := range v {
			fmt.Printf("%v\n, %v\n", l, m)
		}
	}
}
