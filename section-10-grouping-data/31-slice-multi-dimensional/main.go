package main

import "fmt"

func main() {
	iceCreams := []string{"chocolate", "oreo", "ferrero", "chocomint"}
	fmt.Println(iceCreams)
	nieves := []string{"lemon", "lime", "mango", "kiwi"}
	fmt.Println(nieves)

	icn := [][]string{iceCreams, nieves}
	fmt.Println(icn)
}
