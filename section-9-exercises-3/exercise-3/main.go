package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	birthYear := 1986
	currentYear := now.Year()
	yearsLived := -1 // We don't count current year until the year has passed
	for y := birthYear; y < currentYear; y++ {
		fmt.Println(y)
		yearsLived++
	}
	fmt.Printf("Years lived %v\n", yearsLived)
}
