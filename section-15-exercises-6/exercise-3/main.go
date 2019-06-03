package main

import "fmt"

var disclaimerShowed int

func main() {
	showImportantMessage("You need to code your exercises")
	defer importantDisclaimer()
	showImportantMessage("Exercises are good")
	showImportantMessage("Exercises must be completed before submitting")
}

func showImportantMessage(s string) {
	defer importantDisclaimer()
	fmt.Println(s)
}

func importantDisclaimer() {
	disclaimerShowed++
	fmt.Println("All of the rights of this program are reserved", disclaimerShowed)
}
