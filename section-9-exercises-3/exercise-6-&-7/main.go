package main

import (
	"fmt"
	"time"
)

func main() {
	modulusNum := 3
	for x := 0; x <= time.Now().Second(); x++ {
		if x%modulusNum == 0 {
			fmt.Printf("Modulus is 0 for: %v\n", x)
		} else if x%modulusNum == 1 {
			fmt.Printf("Modulus is 1, for: %v\n", x)
		} else {
			fmt.Printf("Modulus is %v, for: %v\n", x%modulusNum, x)
		}
	}
}
