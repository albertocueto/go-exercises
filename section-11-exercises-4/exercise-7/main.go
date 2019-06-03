package main

import "fmt"

func main() {
	slice1 := []string{"Dalinar", "Renarin", "Adolin", "Evi"}
	slice2 := []string{"Kaladin", "Shallan", "Jazhna", "Sylfrina"}
	y := [][]string{slice1, slice2}
	fmt.Println(y)
	for _, v := range y {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
