package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m := map[string][]string{
		"cueto_jose":   {"videogames", "reading", "rock climbing", "programming", "movies"},
		"arvizu_diana": {"Intesive care", "Learning", "Biking", "Traveling"},
	}
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
		for _, v2 := range m[k] {
			fmt.Printf("%v\n", v2)
		}
	}
	fmt.Println("After adding another person")
	m["cueto_miguel"] = []string{"videogames", "comics", "movies"}
	keys := []string{}
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
		keys = append(keys, k)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	rn := r1.Intn(3)
	keyToBeDeleted := keys[rn]

	delete(m, keyToBeDeleted)
	fmt.Println("After deleting one random record: ", keyToBeDeleted)

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
		keys = append(keys, k)
	}
}
