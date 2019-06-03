package main

import (
	"fmt"
	"time"
)

func main() {
	currentYear := time.Now()
	birthYear := 1986
	x := birthYear
	yearsLived := -1
	for {
		if x > currentYear.Year() {
			break
		}
		fmt.Println(x)
		x++
		yearsLived++
	}
	fmt.Printf("Years lived: %v\n", yearsLived)
}
