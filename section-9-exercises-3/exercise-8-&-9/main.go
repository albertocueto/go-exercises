package main

import "fmt"

func main() {
	switch {
	case 2 == 2:
		fmt.Println("2 is certainly equal to 2")
		fallthrough
	case 4 == 2:
		fmt.Println("4 isn't 2, but we fell through...")
	case "unreachable" == "reachable":
		fmt.Println("This code can never be reached, it's always false, should delete it")
	default:
		fmt.Println("Should be reached if everything else's a lie")
	}

	favSport := "Rock climbing"
	switch favSport {
	case "Rock climbing":
		fmt.Println("There is nothing better than rock climbing man")
		fallthrough
	case "Biking":
		fmt.Println("Although... Biking is pretty neat")
	case "Swimming":
		fmt.Println("I like it, but makes me nervous")
	case "Soccer":
		fmt.Println("Not for me")
	default:
		fmt.Println("I'm a potato")
	}
}
