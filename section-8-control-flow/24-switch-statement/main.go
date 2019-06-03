package main

import "fmt"

func main() {
	//fallthrough makes the control go to the next case even though the case statement isn't true
	switch {
	case false:
		fmt.Println("This shouldn't print")
	case (2 == 4):
		fmt.Println("This shouldn't print")
	case (3 == 3):
		fmt.Println("This should print 3")
		fallthrough
	case (4 == 4):
		fmt.Println("This might maybe print by chance 4")
		fallthrough
	case (7 == 9):
		fmt.Println("This might maybe print by chance 9")
		fallthrough
	case (15 == 15):
		fmt.Println("This might maybe print by chance 15")
		fallthrough
	default:
		fmt.Println("This is the default print")
	}

	character := "Kaladin"
	switch character {
	case "Kaladin":
		fmt.Println("Stormblessed")
	case "Dalinar", "Renarin", "Adolin":
		fmt.Println("Reign with honor, Kholin family")
	case "Shallan":
		fmt.Println("The best version of herself")
	default:
		fmt.Println("Any other character")
	}

	// Switch expressions and case expressions are evaluated
	// If the switch expression is empty, that is always taken as a true
	// Remember that fallthrough is not automatic
	switch {
	case true:
		fmt.Println("Always true")
	case false:
		fmt.Println("Never false")
	}
}
